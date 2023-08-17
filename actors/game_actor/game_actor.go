package game_actor

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/go-redis/redis"
	"go_server/actors"
	"go_server/common"
	"net"
)

type GameActor struct {
	listener *net.Listener
	conn     *redis.Client
	fn       func(context actor.Context) bool
}

func (this *GameActor) Start() error {
	return nil
}

func (this *GameActor) Receive(context actor.Context) {
	if this.fn != nil {
		if this.fn(context) {
			return
		}
	}

	switch msg := context.Message().(type) {
	case *actors.DBStartReq:
		this.Start()
	case *common.HotfixReq:
		fn, initfn := common.GetSoFun(msg.Path, msg.FnName, msg.InitName)
		if initfn(this) {
			this.fn = fn
		}
	}
}
