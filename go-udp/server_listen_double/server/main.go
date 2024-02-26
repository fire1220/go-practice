package main

import (
	"log"
	"net"
	"time"
)

func main() {
	local := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 2000}
	conn, err := net.ListenUDP("udp", local)
	if err != nil {
		log.Printf("链接失败：err=%v\n", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("关闭连接失败；err=%v\n", err)
		}
	}()
	for {
		data := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("接收客户端数据失败；err=%v\n", err)
			continue
		}
		log.Printf("成功接收客户端%v个字节，客户端地址：%v，内容是：%v\n", n, addr, string(data))
		go send(conn, addr, []byte("服务器已经收到消息了"))
	}
}

func send(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
	time.Sleep(1 * time.Second)
	n, err := conn.WriteToUDP(data, addr)
	if err != nil {
		log.Printf("发送数据失败；err:%v\n", err)
	}
	log.Printf("成功发送%v个字节，目标地址：%v，内容：%v\n", n, addr.String(), string(data))
}
