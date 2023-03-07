package actiongame

import (
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

type State byte

const (
	STATE_JOIN State = iota
	STATE_LEAVE
)

type PlayerConn struct {
	id    Uid
	state State
}

func (c *PlayerConn) ToPB() *pb.UserConn {
	return &pb.UserConn{Id: uint32(c.id)}
}

var (
	world = &World{
		room: pubsub.New(),
	}
)

func GetWorld( /* id */ ) *World {
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
	conn := &PlayerConn{id, STATE_JOIN}
	world.room.Pub(conn)
	return world, conn
}

func Leave(id Uid) {
	delete(world.players, id)
	world.room.Pub(&PlayerConn{id, STATE_LEAVE})
}
