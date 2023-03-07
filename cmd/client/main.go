package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "actiongame/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "The server address")
)

func init() { flag.Parse() }

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("did not connect: ", err)
	}
	defer conn.Close()
	c := pb.NewGameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Hello(ctx, &pb.HelloReq{Name: "taro"})
	if err != nil {
		log.Panic("could not greet: ", err)
	}
	log.Printf("Greeting: %s", r.GetMsg())
}
