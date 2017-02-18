package etcdclient

import (
	log "github.com/Sirupsen/logrus"
	etcdClient "github.com/coreos/etcd/client"
	"os"
	"strings"
)

const (
	DEFAULT_ETCD = "http://127.0.0.1:2379"
)

var endpoints []string
var client etcdClient.Client

func init() {
	// etcd client
	endpoints = []string{DEFAULT_ETCD}
	if env := os.Getenv("ETCD_HOST"); env != "" {
		endpoints = strings.Split(env, ";")
	}

	// config
	cfg := etcdClient.Config{
		Endpoints: endpoints,
		Transport: etcdClient.DefaultTransport,
	}

	// create client
	c, err := etcdClient.New(cfg)
	if err != nil {
		log.Error(err)
		return
	}
	client = c
}

func KeysAPI() etcdClient.KeysAPI {
	return etcdClient.NewKeysAPI(client)
}
