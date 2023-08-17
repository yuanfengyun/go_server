package actors

import (
	"encoding/json"
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/go-redis/redis"
	"go_server/common"
	"net"
)

type DBStartReq struct {
}

type DBLoadAccountReq struct {
	Account string
}

type DBLoadAccountRsp struct {
	err error
	acc common.Account
}

type DBCreatAccountReq struct {
	Account string
}

type DBCreatAccountRsp struct {
	err error
	acc common.Account
}

type DBActor struct {
	listener *net.Listener
	conn     *redis.Client
	fn       func(context actor.Context) bool
}

func (this *DBActor) Start() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	this.conn = rdb
	return nil
}

func (this *DBActor) Receive(context actor.Context) {
	if this.fn != nil {
		if this.fn(context) {
			return
		}
	}

	switch msg := context.Message().(type) {
	case *DBStartReq:
		this.Start()
	case *common.HotfixReq:
		fn, initfn := common.GetSoFun(msg.Path, msg.FnName, msg.InitName)
		if initfn(this) {
			this.fn = fn
		}
	case *DBLoadAccountReq:
		fmt.Println("handle DBLoadAccountReq")
		val, err := this.conn.Get("account_" + msg.Account).Result()
		rsp := &DBLoadAccountRsp{
			err: err,
		}
		if err == nil {
			err = json.Unmarshal([]byte(val), &rsp.acc)
			if err != nil {
				rsp.err = err
			}
		}
		context.Respond(rsp)
	}
}
