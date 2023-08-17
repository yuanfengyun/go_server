package common

import (
	"github.com/asynkron/protoactor-go/actor"
	"plugin"
)

type HotfixReq struct {
	Path     string
	FnName   string
	InitName string
}

func GetSoFun(path string, name string, initname string) (func(context actor.Context) bool, func(interface{}) bool) {
	so, err := plugin.Open(path)
	if err != nil {
		return nil, nil
	}

	fn, err1 := so.Lookup(name)
	if err1 != nil {
		return nil, nil
	}

	f := fn.(func(context actor.Context) bool)

	fninit, err2 := so.Lookup(initname)
	if err2 != nil {
		return nil, nil
	}

	finit := fninit.(func(interface{}) bool)

	return f, finit
}
