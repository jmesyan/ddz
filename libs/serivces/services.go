package serivces

import (
	"sync"

	etcdClient "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
)

const (
	DEFAULT_ETCD         = "http://127.0.0.1:2379"
	DEFAULT_SERVICE_PATH = "/backends"
	DEFAULT_NAME_FILE    = "backends/names"
)

// a single connection
type client struct {
	key  string
	conn *grpc.ClientConn
}

// service structure
type service struct {
	clients []client
	idx     uint32 // round-robin
}

// all services
type servicePool struct {
	sync.RWMutex
	services         map[string]*service
	knownNames       map[string]bool // store names.txt
	nameCheckEnabled bool
	client           etcdClient.Client
	callbacks        map[string][]chan string // service add callback notify
}

var (
	_defaultPool servicePool
	once         sync.Once
)

// Init() MUST be called before using
func Init(names ...string) {
	once.Do(func() {
		_defaultPool.init(names...)
	})
}

func (p *servicePool) init(names ...string) {
	// etcd client
	endpoints := []string{DEFAULT_ETCD}
	if env := os.Getenv("ETCD_HOST"); env != "" {
		endpoints = strings.Split(env, ";")
	}

	// init etcd client
	cfg := etcdClient.Config{
		Endpoints: endpoints,
		Transport: etcdClient.DefaultTransport,
	}
	c, err := etcdClient.New(cfg)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	p.client = c

	// init
	p.services = make(map[string]*service)
	p.knownNames = make(map[string]bool)

	// name list
	if len(names) == 0 {
		names = p.loadNames
	}
	if len(names) > 0 {
		p.nameCheckEnabled = true
	}

	log.Println("all service names:", names)
	for _, v := range names {
		p.knownNames[DEFAULT_SERVICE_PATH+"/"+strings.TrimSpace(v)] = true
	}

	// start connection
	p.connectAll(DEFAULT_SERVICE_PATH)
}

func (p *servicePool) newEtcdKeysAPI() etcdClient.KeysAPI {
	return etcdClient.NewKeysAPI(p.client)
}

func (p *servicePool) findService(path string) (service, bool) {
	service := p.services[path]
	if service == nil || len(service.clients) == 0 {
		return nil, false
	}

	return service, true
}

// get stored service name
func (p *servicePool) loadNames() []string {
	keysAPI := p.newEtcdKeysAPI()
	// get the keys under the directory
	log.Println("reading names:", DEFAULT_NAME_FILE)
	resp, err := keysAPI.Get(context.Background(), DEFAULT_NAME_FILE, nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	// validation
	if resp.Node.Dir {
		log.Println("names is not a file")
		return nil
	}

	// split names
	return strings.Split(resp.Node.Value, "\n")
}

// connect to all services
func (p *servicePool) connectAll(directory string) {
	keysAPI := p.newEtcdKeysAPI()
	// get the keys under the directory
	log.Println("connecting the services under:", directory)
	resp, err := keysAPI.Get(context.Background(), directory, &etcdClient.GetOptions{Recursive: true})
	if err != nil {
		log.Println(err)
		return
	}

	// validate
	if !resp.Node.Dir {
		log.Println("not a directory")
		return
	}

	for _, node := range resp.Node.Nodes {
		if node.Dir {
			for _, service := range node.Nodes {
				p.addService(service.Key, service.Value)
			}
		}
	}
	log.Println("services add complete")

	go p.watcher()
}

// watcher for data change in etcd director
func (p *servicePool) watcher() {
	keysAPI := p.newEtcdKeysAPI()
	w := keysAPI.Watcher(DEFAULT_SERVICE_PATH, &etcdClient.WatcherOptions{Recursive: true})
	for {
		resp, err := w.Next(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.Node.Dir {
			continue
		}

		switch resp.Action {
		case "set", "create", "update", "compareAndSwap":
			p.addService(resp.Node.Key, resp.Node.Value)
		case "delete":
			p.removeService(resp.PrevNode.Key)
		}
	}
}

// add a service
func (p *servicePool) addService(key, value string) {
	p.Lock()
	defer p.Unlock()
	serviceName := filepath.Dir(key)
	if p.nameCheckEnabled && !p.knownNames[serviceName] {
		return
	}

	// try new service kind init
	if p.services[serviceName] != nil {
		p.services[serviceName] = &service{}
	}

	// create service connection
	service := p.services[serviceName]
	if conn, err := grpc.Dial(value, grpc.WithBlock(), grpc.WithInsecure()); err == nil {
		service.clients = append(service.clients, client{key, conn})
		log.Println("service added:", key, "-->", value)
		for k := range p.callbacks[serviceName] {
			select {
			case p.callbacks[serviceName][k] <- key:
			default:
			}
		}
	} else {
		log.Println("did not connect:", key, "-->", value, "error:", err)
	}
}

// remove a service
func (p *servicePool) removeService(key string) {
	p.Lock()
	defer p.Unlock()
	// name check
	serviceName := filepath.Dir(key)
	if p.nameCheckEnabled && !p.knownNames[serviceName] {
		return
	}

	// check service kind
	service := p.services[serviceName]
	if service == nil {
		log.Println("no such service:", serviceName)
		return
	}

	// remove service
	for k := range service.clients {
		if service.clients[k].key == key {
			service.clients[k].conn.Close()
			service.clients = append(service.clients[:k], service.clients[k+1:]...)
			log.Println("service removed:", key)
			return
		}
	}
}

// provide a specific key for a service, eg:
// path:/backends/snowflake, id:s1
//
// the full canonical path for this service is:
// /backends/snowflake/s1
func (p *servicePool) getServiceWithId(path string, id string) *grpc.ClientConn {
	p.RLock()
	defer p.RUnlock()

	// find service
	service, exist := p.findService(path)
	if !exist {
		return nil
	}

	// find a service with specific id
	fullPath := string(path) + "/" + id
	for k := range service.clients {
		if service.clients[k].key == fullPath {
			return service.clients[k].conn
		}
	}

	return nil
}

// get a service in round-robin style
// especially useful for load-balance with state-less services
func (p *servicePool) getService(path string) (conn *grpc.ClientConn, key string) {
	p.RLock()
	defer p.RUnlock()

	// find service
	service, exist := p.findService(path)
	if !exist {
		return nil, ""
	}

	idx := int(atomic.AddUint32(&service.idx, 1)) % len(service.clients)
	return service.clients[idx].conn, service.clients[idx].key
}

func (p *servicePool) registerCallback(path string, callback chan string) {
	p.Lock()
	defer p.Unlock()
	if p.callbacks == nil {
		p.callbacks = make(map[string][]chan string)
	}

	p.callbacks[path] = append(p.callbacks[path], callback)
	if s, ok := p.services[path]; ok {
		for k := range s.clients {
			callback <- s.clients[k].key
		}
	}
	log.Println("register callback on:", path)
}

/////////////////////////////////////////////////////////////////
// Wrappers
func GetService(path string) *grpc.ClientConn {
	conn, _ := _defaultPool.getService(path)
	return conn
}

func GetService2(path string) (*grpc.ClientConn, string) {
	conn, key := _defaultPool.getService(path)
	return conn, key
}

func GetServiceWithId(path string, id string) *grpc.ClientConn {
	return _defaultPool.getServiceWithId(path, id)
}

func RegisterCallback(path string, callback chan string) {
	_defaultPool.registerCallback(path, callback)
}
