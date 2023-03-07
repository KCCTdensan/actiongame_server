package actiongame

import (
	"context"
	"fmt"

	pb "actiongame/proto"
)

type GameServer struct {
	pb.UnimplementedGameServer
}

func (*GameServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRes, error) {
	msg := fmt.Sprintf("Hello, %v!", req.Name)
	return &pb.HelloRes{Msg: msg}, nil
}

func (*GameServer) JoinSession(ctx context.Context, req *pb.JoinReq) (*pb.GameSession, error) {
	return &pb.GameSession{Id: "global"}, nil
}
