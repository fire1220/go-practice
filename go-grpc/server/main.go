package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	pb "main/proto/model"
	"net"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) GetUserInfo(ctx context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("有客户端调用GetUserInfo方法。昵称是：%v\n", r.GetName())
	if len(r.GetName()) == 0 {
		return nil, errors.New("the query name cannot be empty")
	}
	return &pb.UserResponse{Id: 123, Username: "jock", Nickname: r.GetName()}, nil
}

func main() {
	log.Printf("启动服务端")
	// 监听端口
	ln, err := net.Listen("tcp", "127.0.0.1:9888")
	if err != nil {
		log.Printf("listen err: %v\n", err)
		return
	}
	// 实例化gRPC服务
	server := grpc.NewServer()
	// 注册UserService服务
	pb.RegisterUserServiceServer(server, new(UserService))
	// 向gRPC服务端注册反射服务
	reflection.Register(server)
	// 启动gRPC服务
	if err := server.Serve(ln); err != nil {
		log.Printf("启动gRPC服务失败:%v\n", err)
	}
}
