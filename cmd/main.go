package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"actiongame"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func init() { flag.Parse() }

func main() {
	addr := fmt.Sprintf(":%d", *port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	actiongame.Serve(ln)
}
