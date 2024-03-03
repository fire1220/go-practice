package main

import (
	"errors"
	"log"
	"main/proto"
	"net"
	"net/rpc"
)

type UserHandler struct {
}

func (u *UserHandler) GetUserInfo(name string, p *proto.UserInfo) error {
	log.Printf("服务器参数:%v\n", name)
	if name == "" {
		return errors.New("参数不能为空")
	}
	p.Id = 1
	p.Name = name
	p.Nickname = "法外狂徒"
	return nil
}

func main() {
	err := rpc.Register(new(UserHandler))
	if err != nil {
		log.Printf("服务器注册服务失败：%v\n", err)
		return
	}
	listen, err := net.Listen("tcp", "127.0.0.1:9888")
	if err != nil {
		log.Printf("无服务监听失败：%v\n", err)
		return
	}
	log.Printf("启动服务器")
	rpc.Accept(listen)
}
