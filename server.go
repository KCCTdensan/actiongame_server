package actiongame

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"google.golang.org/grpc/metadata"

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

	// todo: まともにする
	room := make(chan *PlayerConn)
	world.room.Sub(func(conn *PlayerConn) { room <- conn })
	<-room // ひどい

	for {
		select {
		case conn := <-room:
			stream.Send(conn.ToPB())
		case <-stream.Context().Done():
			Leave(conn.id)
			return nil
		}
	}
}

func (*GameServer) Move(stream pb.Game_MoveServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("no metadata")
	}
	if len(md["id"]) == 0 {
		return errors.New("no id")
	}

	uid, err := strconv.Atoi(md["id"][0])
	if err != nil {
		return err
	}

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
			world.room.Pub(&pb.UserMove{Id: uint32(uid), Pos: pos})
			fmt.Println(uid, pos) /////////////////////////////////////////
		}
	}()

	// send to client
	world.room.Sub(func(move *pb.UserMove) {
		stream.Send(move)
	})

	<-stream.Context().Done()
	return nil
}
