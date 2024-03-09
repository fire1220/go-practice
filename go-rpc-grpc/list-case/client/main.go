package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pd "main/porot/model"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("客户端启动失败%v\n", err)
		return
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	param := pd.UserFind{Name: "李四", Ids: []int32{1, 2}}
	res, err := pd.NewUserClient(conn).GetUserList(ctx, &param)
	if err != nil {
		log.Printf("rpc请求失败%v\n", err)
		return
	}
	log.Printf("结果：%+v\n", res)
}
