package main

import (
	"fmt"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"time"
)

func main() {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(func() actor.Actor { return &GateActor{} })

	pid := system.Root.Spawn(props)
	//system.Root.Send(pid, &EchoReq{Who: "Roger"})

	result, _ := system.Root.RequestFuture(pid, &EchoReq{Who: "Roger"}, 30*time.Second).Result() // await result

	fmt.Println(result)
	fmt.Println("over")
	_, _ = console.ReadLine()
}
