package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "main/proto/model"
	"time"
)

func main() {
	// credentials, _ := credentials.NewClientTLSFromFile("./ca.crt", "") // (如果服务器配置了ca证书加密)ca证书,第二个参数用来验证，一班是空字符串
	// conn, err := grpc.Dial("127.0.0.1:9888", grpc.WithTransportCredentials(credentials))
	// 远程连接凭证，insecure模式下禁用了传输安全认证
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // 默认客户端必须加密传输，这里禁用加密传输
	conn, err := grpc.Dial("127.0.0.1:9888", credentials)
	if err != nil {
		log.Printf("连接失败：%v\n", err)
		return
	}
	defer conn.Close()
	// 初始化UserService客户端
	userClient := pb.NewUserServiceClient(conn)
	// 定义超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 远程调用服务端方法
	response, err := userClient.GetUserInfo(ctx, &pb.UserRequest{Name: "fire"})
	if err != nil {
		log.Printf("调用GetUserInfo失败:%v\n", err)
		return
	}
	log.Printf("%+v\n", response)
}
