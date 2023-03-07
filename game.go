package actiongame

import (
	"fmt"
	"math/rand"

	"github.com/mattn/go-pubsub"

	pb "actiongame/proto"
)

type Uid uint32

type World struct {
	players map[Uid]*Player
	room    *pubsub.PubSub
}

type Player struct {
	name string
}

type PlayerConn struct {
	id    Uid
	state pb.ConnState
}

func (c *PlayerConn) ToPB() *pb.UserConn {
	return &pb.UserConn{
		Type: pb.ConnState(c.state),
		Id:   uint32(c.id),
	}
}

var (
	world = &World{
		players: make(map[Uid]*Player),
		room:    pubsub.New(),
	}
)

func GetWorld() *World {
	return world
}

func Join(name string) (*World, *PlayerConn) {
	// todo: いい感じにする
	var id Uid
	for {
		id = Uid(rand.Uint32())
		if _, ok := world.players[id]; !ok {
			break
		}
	}
	world.players[id] = &Player{name}
	conn := &PlayerConn{id, pb.ConnState_JOIN}
	fmt.Println("joined", name, "<>", world.players)
	world.room.Pub(conn)
	return world, conn // connはserverstream用
}

func Leave(id Uid) {
	name := world.players[id].name
	delete(world.players, id)
	fmt.Println("leave", name, "<>", world.players)
	world.room.Pub(&PlayerConn{id, pb.ConnState_LEAVE})
}
