package main

import (
	log "github.com/Sirupsen/logrus"
	pb "github.com/master-g/golandlord/snowflake/proto"
	"google.golang.org/grpc"
	"net"
	"os"
)

const (
	_port = ":40001"
)

func main() {
	// listen
	listen, err := net.Listen("tcp", _port)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	log.Info("listening on ", listen.Addr())

	// register service
	s := grpc.NewServer()
	instance := &server{}
	instance.init()
	pb.RegisterSnowflakeServiceServer(s, instance)

	// Start service
	s.Serve(listen)
}
