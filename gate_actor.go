package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	pp "go_server/proto"
	"io"
	"log"
	"net"
)

type Client struct {
	Status    string //"init" "login" "role" "create_role" "game" "closing"
	Conn      net.Conn
	SendQueue []byte
}

type EchoReq struct {
	Who string
}

type EchoRsp struct {
	ret string
}

type GateActor struct {
	listener *net.Listener
}

func (this *GateActor) ListenAndServe(address string) {
	// 绑定监听地址
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err: %v", err))
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
		this.listener = nil
	}(listener)

	this.listener = &listener

	log.Println(fmt.Sprintf("bind: %s, start listening...", address))

	go func() {
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
	headbuf := make([]byte, 8)
	buff := make([]byte, 1024)
	r := bufio.NewReader(conn)
	for {
		_, err := io.ReadFull(r, headbuf)
		if err != nil {
			break
		}
		msglen := binary.BigEndian.Uint32(headbuf)
		id := binary.BigEndian.Uint32(headbuf[4:])
		body := buff[:msglen-4]
		_, err = io.ReadFull(r, body)
		if err != nil {
			break
		}
		err1, msg := pp.Decode(id, body)
		if err1 != nil {
			return
		}
		fmt.Println(msg)
	}

}

func (this *GateActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *EchoReq:
		fmt.Printf("hello %v", msg.Who)
		context.Respond("hi")
		this.ListenAndServe("0.0.0.0:6000")
	}
}
