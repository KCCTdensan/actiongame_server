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
	id   Uid
	name string
}

type PlayerConn struct {
	id    Uid
	state pb.ConnState
}

func (c *PlayerConn) ToPB() *pb.PlayerConn {
	return &pb.PlayerConn{
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

func GenId() (id Uid) {
	for {
		id = Uid(rand.Uint32())
		if _, ok := world.players[id]; !ok {
			return
		}
	}
}

func Register(name string) *Player {
	id := GenId()
	player := &Player{id, name}
	world.players[id] = player
	return player
}

func Join(player *Player) *PlayerConn {
	// todo: いい感じにする
	conn := &PlayerConn{player.id, pb.ConnState_JOIN}
	fmt.Println("joined", player.name, "<>", world.players)
	world.room.Pub(conn)
	return conn
}

func Leave(id Uid) {
	player := world.players[id]
	delete(world.players, id)
	fmt.Println("leave", player.name, "<>", world.players)
	world.room.Pub(&PlayerConn{id, pb.ConnState_LEAVE})
}
