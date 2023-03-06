package actiongame

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "actiongame/proto"
)

type server struct {
	pb.UnimplementedGameServer
}

func (s *server) Hello(ctx context.Context, in *pb.HelloReq) (*pb.HelloRes, error) {
	msg := fmt.Sprintf("hello, %v", in.Name)
	return &pb.HelloRes{Msg: msg}, nil
}

func (s *server) JoinSession(ctx context.Context, in *pb.JoinReq) (*pb.GameSession, error) {
	return &pb.GameSession{Id: "global"}, nil
}

func Serve(ln net.Listener) {
	s := grpc.NewServer()
	pb.RegisterGameServer(s, &server{})

	log.Print("server listening at:", ln.Addr())
	if err := s.Serve(ln); err != nil {
		log.Panic("failed to serve:", err)
	}
}
