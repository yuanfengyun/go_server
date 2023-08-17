package actors

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/golang/protobuf/proto"
	msg "go_server/proto"
	"io"
	"log"
	"net"
	"time"
)

type Client struct {
	Status    string //"init" "login" "role" "create_role" "game" "closing"
	Conn      net.Conn
	OldConn   net.Conn
	SendQueue []byte
	Account   string
}

var client_map map[string]*Client = map[string]*Client{}

type GateActor struct {
	listener *net.Listener
}

func (this *GateActor) Start(address string) {
	this.ListenAndServe(address)
}

func (this *GateActor) ListenAndServe(address string) {
	// 绑定监听地址
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err: %v", err))
		return
	}

	this.listener = &listener

	log.Println(fmt.Sprintf("bind: %s, start listening...", address))

	go func() {
		defer func(listener net.Listener) {
			err := listener.Close()
			if err != nil {

			}
			this.listener = nil
		}(listener)

		for {
			// Accept 会一直阻塞直到有新的连接建立或者listen中断才会返回
			conn, err := listener.Accept()
			if err != nil {
				// 通常是由于listener被关闭无法继续监听导致的错误
				log.Fatal(fmt.Sprintf("accept err: %v", err))
			}
			// 开启新的 goroutine 处理该连接
			go this.HandleConn(conn)
		}
	}()
}

func (this *GateActor) HandleConn(conn net.Conn) {
	headbuf := make([]byte, 4)
	buff := make([]byte, 1024)
	r := bufio.NewReader(conn)
	//var client *Client = nil
	Status := "init"
	for {
		_, err := io.ReadFull(r, headbuf)
		if err != nil {
			break
		}
		msglen := binary.BigEndian.Uint32(headbuf)
		body := buff[:msglen]
		_, err = io.ReadFull(r, body)
		if err != nil {
			break
		}
		req := &msg.Req{}
		err1 := proto.Unmarshal(body, req)
		if err1 != nil {
			break
		}
		if Status == "init" {
			if req.Command != "login" {
				break
			}
			Status = "login"

			// 向登陆服发出请求
			result, _ := System.Root.RequestFuture(
				login_actor,
				&LoginReq{Account: req.StringParams[0], Passwd: req.StringParams[1]},
				10*time.Second).Result()

			LoginRsp := result.(*LoginRsp)
			if LoginRsp.err != nil {
				break
			}

			// 检查是否已经登陆

			// 登陆进游戏
		} else if Status == "login" {
			continue
		}

		fmt.Println(req)
	}
}

func (this *GateActor) HandleMsg(conn net.Conn, msg *msg.Req) {

}

func (this *GateActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *CommandReq:
		command := msg.Command
		if command == "start" {
			this.Start(msg.Params[0])
		}
	}
}
