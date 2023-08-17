package actors

import (
	"github.com/asynkron/protoactor-go/actor"
	"go_server/common"
)

type LoginReq struct {
	Account string
	Passwd  string
}

type LoginRsp struct {
	err     error
	Account common.Account
}

type LoginActor struct {
	fn func(context actor.Context) bool
}

func (this *LoginActor) Start() error {
	return nil
}

func (this *LoginActor) Receive(context actor.Context) {
	if this.fn != nil {
		if this.fn(context) {
			return
		}
	}

	switch msg := context.Message().(type) {
	case *CommandReq:
		this.Start()
	case *common.HotfixReq:
		fn, initfn := common.GetSoFun(msg.Path, msg.FnName, msg.InitName)
		if initfn(this) {
			this.fn = fn
		}
	}
}
