package actiongame

import (
	"context"
	"fmt"
	"io"

	pb "actiongame/proto"
)

type GameServer struct {
	pb.UnimplementedGameServer
}

func (*GameServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRes, error) {
	msg := fmt.Sprintf("Hello, %v!", req.Name)
	return &pb.HelloRes{Msg: msg}, nil
}

func (*GameServer) Join(req *pb.JoinReq, stream pb.Game_JoinServer) error {
	world, conn := Join(req.Name)
	stream.Send(conn.ToPB())

	room := make(chan *PlayerConn)
	world.room.Sub(func(conn *PlayerConn) { room <- conn })

	for {
		select {
		case <-stream.Context().Done():
			Leave(conn.id)
			return nil
		case conn := <-room:
			stream.Send(conn.ToPB())
		}
	}
}

func (*GameServer) Move(stream pb.Game_MoveServer) error {
	// get user id from context
	uid := uint32(0) //
	world := GetWorld()

	// recv from client
	go func() {
		for {
			pos, err := stream.Recv()
			if err == io.EOF {
				break // nil
			}
			if err != nil {
				break // err
			}
			world.room.Pub(pos)
		}
		// close connection
	}()

	// send to client
	world.room.Sub(func(pos *pb.Pos) {
		for {
			stream.Send(&pb.UserMove{Id: uid, Pos: pos})
		}
	})
}
