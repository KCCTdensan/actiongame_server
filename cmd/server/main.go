package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"actiongame"
	pb "actiongame/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func init() { flag.Parse() }

func main() {
	addr := fmt.Sprintf(":%d", *port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic("failed to listen: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterGameServer(s, &actiongame.GameServer{})

	log.Print("server listening at: ", ln.Addr())
	if err := s.Serve(ln); err != nil {
		log.Panic("failed to serve: ", err)
	}
}
