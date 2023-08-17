package main

import (
	"fmt"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"go_server/actors"
	game_actor "go_server/actors/game_actor"
	"go_server/common"
)

func main() {
	system := actor.NewActorSystem()
	// 启动网关actor
	actors.StartGateActor("0.0.0.0:8000")

	fmt.Println("[info] gate started")

	actors.StartDBActor()
	fmt.Println("[info] db started")

	actors.StartLoginActor()
	fmt.Println("[info] login started")

	game_actor.Start()
	fmt.Println("[info] db started")

	req1 := &actors.DBLoadAccountReq{
		Account: "account",
	}

	// 更新数据库actor
	db := actors.GetActor("db")
	system.Root.Send(db, req1)
	_, _ = console.ReadLine()

	// 加载动态库热更
	req := &common.HotfixReq{
		Path:     "./dbactor.so",
		FnName:   "Fn",
		InitName: "InitFn",
	}
	system.Root.Send(db, req)

	// 测试热更后效果
	system.Root.Send(db, req1)

	_, _ = console.ReadLine()
}
