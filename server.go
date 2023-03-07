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
	world := GetWorld()
	me := Register(req.Name)

	stream.Send(&pb.PlayerConn{
		Id:   uint32(me.id),
		Type: pb.ConnState_JOIN,
	})
	for id := range world.players {
		if id != me.id {
			stream.Send(&pb.PlayerConn{
				Id:   uint32(id),
				Type: pb.ConnState_SPAWN,
			})
		}
	}

	// todo: まともにする
	room := make(chan *PlayerConn)
	world.room.Sub(func(conn *PlayerConn) { room <- conn })

	Join(me)
	<-room // 自分のJOINは無視

	for {
		select {
		case conn := <-room:
			stream.Send(conn.ToPB())
		case <-stream.Context().Done():
			Leave(me.id)
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
			world.room.Pub(&pb.PlayerMove{Id: uint32(uid), Pos: pos})
			fmt.Println(uid, pos) /////////////////////////////////////////
		}
	}()

	// send to client
	world.room.Sub(func(move *pb.PlayerMove) {
		stream.Send(move)
	})

	<-stream.Context().Done()
	return nil
}
