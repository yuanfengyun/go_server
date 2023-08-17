package actors

import "github.com/asynkron/protoactor-go/actor"

type CommandReq struct {
	Command string
	Params  []string
	Index   uint32
}

type CommandRsp struct {
	Command string
	Result  []string
	Index   uint32
}

var System *actor.ActorSystem = nil
var Root *actor.RootContext = nil
var db_actor *actor.PID = nil
var gate_actor *actor.PID = nil
var game_actor *actor.PID = nil
var login_actor *actor.PID = nil

func init() {
	System = actor.NewActorSystem()
	Root = System.Root
}

func StartDBActor() {
	// 启动数据库actor
	props := actor.PropsFromProducer(func() actor.Actor { return &DBActor{} })

	db_actor = System.Root.Spawn(props)

	System.Root.Send(db_actor, &DBStartReq{})
}

func StartGateActor(addr string) {
	// 启动数据库actor
	props := actor.PropsFromProducer(func() actor.Actor { return &GateActor{} })
	gate_actor = System.Root.Spawn(props)

	System.Root.Send(gate_actor, &CommandReq{Command: "start", Params: []string{addr}, Index: 1})
}

func StartLoginActor() {
	props := actor.PropsFromProducer(func() actor.Actor { return &LoginActor{} })
	login_actor = Root.Spawn(props)

	Root.Send(login_actor, &CommandReq{Command: "start"})
}

func Register(name string, pid *actor.PID) {
	if name == "game" {
		game_actor = pid
	}
}

func GetActor(name string) *actor.PID {
	if name == "gate" {
		return gate_actor
	} else if name == "db" {
		return db_actor
	} else if name == "game" {
		return game_actor
	} else if name == "login" {
		return login_actor
	}
	return nil
}
