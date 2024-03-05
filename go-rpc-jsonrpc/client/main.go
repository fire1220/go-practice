package main

import (
	"fmt"
	"log"
	"main/proto"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:9888")
	if err != nil {
		log.Printf("客户端链接失败：%v\n", err)
	}
	defer conn.Close()
	user := new(proto.UserInfo)
	err = conn.Call("UserHandler.GetUserInfo", "张三", user)
	if err != nil {
		log.Printf("rpc调用失败：%v\n", err)
		return
	}
	fmt.Printf("%#v\n", *user)
}
