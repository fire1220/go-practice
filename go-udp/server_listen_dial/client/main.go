package main

import (
	"log"
	"net"
)

func main() {
	// 拨号模式
	serverAddr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 1888,
	}
	localAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 2000,
	}
	conn, err := net.DialUDP("udp", localAddr, serverAddr) // 这种方式是固定本地端口
	// conn, err := net.Dial("udp", "127.0.0.1:1888") // 这种方式也可以,这种方式是随机本地端口
	if err != nil {
		log.Printf("链接失败：err=%v\n", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	n, err := conn.Write([]byte("你好世界"))
	if err != nil {
		log.Printf("发送数据失败；err=%v\n", err)
	}
	log.Printf("发送%v个字节", n)
	data := make([]byte, 1024)
	r, err := conn.Read(data)
	if err != nil {
		log.Printf("接收数据失败；err=%v\n", err)
	}
	log.Printf("接收%v个字节，内容：%v\n", r, string(data))
}
