package main

import (
	"fmt"
	"log"
	"main/proto"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:9888")
	if err != nil {
		log.Printf("链接失败：%v\n", err)
		return
	}
	defer conn.Close()
	user := new(proto.UserInfo)
	err = conn.Call("UserHandler.GetUserInfo", "李四", user)
	if err != nil {
		log.Printf("远程访问失败%v\n", err)
		return
	}
	fmt.Printf("%#v\n", user)
}
