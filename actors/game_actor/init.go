package game_actor

import (
	"github.com/asynkron/protoactor-go/actor"
	"go_server/actors"
)

var RoleMgr map[string]*Role = map[string]*Role{}

func Start() {
	props := actor.PropsFromProducer(func() actor.Actor { return &GameActor{} })

	game_actor := actors.Root.Spawn(props)

	actors.Root.Send(game_actor, &actors.CommandReq{Command: "start"})

	actors.Register("game", game_actor)
}

func init() {

}
