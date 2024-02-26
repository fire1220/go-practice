package main

import (
	"context"
	"log"
	"net"
)

func main() {
	local := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1888}
	conn, err := net.ListenUDP("udp", local)
	if err != nil {
		log.Printf("创建链接失败:%v\n", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	for {
		data := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("获取数据失败：%v\n", err)
			continue
		}
		log.Printf("接收%v个字节,客户端地址：%v,数据是：%v\n", n, addr, string(data))
		go send(context.Background(), conn, addr, []byte("hello"))
	}

}

func send(ctx context.Context, conn *net.UDPConn, remote *net.UDPAddr, data []byte) {
	n, err := conn.WriteToUDP(data, remote)
	if err != nil {
		log.Printf("发送数据失败：err=%v\n", err)
	}
	log.Printf("成功发送%v个字节", n)
}
