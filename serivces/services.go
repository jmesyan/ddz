package serivces

import (
    "sync"

    etcdClient "github.com/coreos/etcd/client"
    "google.golang.org/grpc"
)

const (
    DEFAULT_ETCD = "http://127.0.0.1:2379"
    DEFAULT_SERVICE_PATH = "/backends"
    DEFAULT_NAME_FILE = "backends/names"
)

// a single connection
type client struct {
    key string
    conn *grpc.ClientConn
}

// service structure
type service struct {
    clients []client
    idx uint32 // round-robin
}

// all services
type service_pool struct {
    services map[string]*service
    known_names map[string]bool // store names.txt
    enable_name_check bool
    client etcdClient.Client
    callbacks map[string][]chan string // service add callback notify
    sync.RWMutex
}