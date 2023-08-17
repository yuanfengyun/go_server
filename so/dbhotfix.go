package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"go_server/actors"
)

func InitFn(i interface{}) bool {
	a := i.(*actors.DBActor)
	fmt.Println("dbhotfix.init")
	fmt.Println(a)
	return true
}

func Fn(context actor.Context) bool {
	fmt.Println("dbhotfix.fn")

	return true
}
