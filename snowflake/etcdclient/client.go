package etcdclient

import (
	log "github.com/Sirupsen/logrus"
	etcdclient "github.com/coreos/etcd/client"
	"os"
	"strings"
)

const (
	//DEFAULT_ETCD = "http://127.0.0.1:2379"
	DEFAULT_ETCD = "http://192.168.0.10:2379"
)

var endpoints []string
var client etcdclient.Client

func init() {
	// etcd client
	endpoints = []string{DEFAULT_ETCD}
	if env := os.Getenv("ETCD_HOST"); env != "" {
		endpoints = strings.Split(env, ";")
	}

	// config
	cfg := etcdclient.Config{
		Endpoints: endpoints,
		Transport: etcdclient.DefaultTransport,
	}

	// create client
	c, err := etcdclient.New(cfg)
	if err != nil {
		log.Error(err)
		return
	}
	client = c
}

func KeysAPI() etcdclient.KeysAPI {
	return etcdclient.NewKeysAPI(client)
}
