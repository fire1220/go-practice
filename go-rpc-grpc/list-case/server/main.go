package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	pd "main/porot/model"
	"net"
)

type UserServer struct {
	pd.UserServer
}

func (u *UserServer) GetUserList(ctx context.Context, find *pd.UserFind) (*pd.UserList, error) {
	if len(find.Ids) < 0 {
		return nil, errors.New("参数不能为空")
	}
	log.Printf("%v\n", find.String())
	find.GetName()
	ret := new(pd.UserList)
	list := make([]*pd.UserInfo, 0, len(find.Ids))
	for _, id := range find.Ids {
		list = append(list, &pd.UserInfo{
			Id:       id,
			Name:     find.GetName(),
			Nickname: "法外狂徒",
		})
	}
	ret.List = list
	return ret, nil
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:9888")
	if err != nil {
		log.Printf("服务器监听失败：%v\n", err)
		return
	}
	server := grpc.NewServer()
	pd.RegisterUserServer(server, new(UserServer))
	reflection.Register(server)
	log.Println("启动服务")
	if err := server.Serve(ln); err != nil {
		log.Printf("服务器启动失败%v\n", err)
	}
}
