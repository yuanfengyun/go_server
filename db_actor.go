package main

import (
	"encoding/json"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/go-redis/redis"
	"net"
)

type DBLoadAccountReq struct {
	Account string
}

type DBLoadAccountRsp struct {
	err error
	acc Account
}

type DBCreatAccountReq struct {
	Account string
}

type DBCreatAccountRsp struct {
	err error
	acc Account
}

type DBActor struct {
	listener *net.Listener
	conn     *redis.Client
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
	switch msg := context.Message().(type) {
	case *DBLoadAccountReq:
		val, err := this.conn.Get("account_" + DBLoadAccountReq.Account).Result()
		rsp := &DBLoadAccountRsp{
			err: err,
		}
		if err == nil {
			err = json.Unmarshal([]byte(val), &rsp.acc)
			if err != nil {

			}
		}
		context.Respond(rsp)
	}
}
