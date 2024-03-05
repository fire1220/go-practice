package main

import (
	"errors"
	"log"
	"main/proto"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type UserHandler struct {
}

func (u *UserHandler) GetUserInfo(name string, reply *proto.UserInfo) error {
	log.Printf("服务器入参：%v\n", name)
	if name == "" {
		return errors.New("参数不能为空")
	}
	reply.Id = 1
	reply.Name = "jock"
	reply.Nickname = "fire"
	return nil

}

func main() {
	err := rpc.Register(new(UserHandler))
	if err != nil {
		log.Printf("服务器注册服务失败；%v\n", err)
		return
	}
	listen, err := net.Listen("tcp", "127.0.0.1:9888")
	if err != nil {
		log.Printf("服务器启动端口监听失败：%v\n", err)
		return
	}
	defer listen.Close()
	log.Printf("服务端启动服务\n")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("服务链接失败：%v\n", err)
			continue
		}
		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}
