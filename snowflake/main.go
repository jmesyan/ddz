/*
 * This is a standard gRPC server implementation
 * the Snowflake.ServiceServer is implement in service.go
 */
package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	pb "github.com/master-g/golandlord/snowflake/proto"
	"google.golang.org/grpc"
	"gopkg.in/urfave/cli.v2"
	"net"
	"os"
	"sort"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Value:       0,
				Aliases:     []string{"p"},
				Usage:       "local port to listen",
				DefaultText: "random",
			},
			&cli.StringFlag{
				Name:    "etcd",
				Value:   "",
				Aliases: []string{"e"},
				Usage:   "etcd server address, if there are multiple servers, use ';' to separate",
			},
		},
		Name:    "snowflake",
		Usage:   "Twitter's UUID generator snowflake in golang",
		Version: "v1.0.0",
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			startSnowflake(port)
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	app.Run(os.Args)
}

func startSnowflake(port int) {
	// listen
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	log.Info("listening on ", lis.Addr())

	// register service
	s := grpc.NewServer()
	instance := &server{}
	instance.init()
	pb.RegisterSnowflakeServiceServer(s, instance)

	// Start service
	s.Serve(lis)
}
