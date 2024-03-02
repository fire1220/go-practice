package main

import (
	"log"
	"net"
	"time"
)

func main() {
	localAddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 3000}
	serverAddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 2000}
	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Printf("客户端监听失败；err=%v", err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("关闭监听失败；err=%v\n", err)
		}
	}()
	n, err := conn.WriteToUDP([]byte("你好世界"), serverAddr)
	if err != nil {
		log.Printf("向服务端发送数据失败；err=%v\n", err)
		return
	}
	log.Printf("成功发送%v个字节", n)
	go listen(conn)

	time.Sleep(4 * time.Second)
}

func listen(conn *net.UDPConn) {
	for {
		data := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("客户端读取数据失败；err=%v\n", err)
		}
		log.Printf("成功接收服务端%v个字节，服务端地址：%v，内容是：%v\n", n, addr, string(data))
	}
}
