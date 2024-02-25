package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:1888")
	if err != nil {
		fmt.Printf("启动监听失败，错误：%v\n", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept() // 如果没有数据接收会阻塞在这里
		if err != nil {
			fmt.Printf("接收链接失败，错误：%v\n", err)
			conn.Close()
			continue
		}
		fmt.Printf("连接成功，来自%v\n", conn.RemoteAddr().String())
		go readWrite(conn)
	}
}

func readWrite(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Printf("关闭链接失败，err:%v\n", err)
		}
	}()
	for {
		readData := make([]byte, 1024)
		n, err := conn.Read(readData)
		if err != nil && err != io.EOF {
			fmt.Printf("读取失败，err:%v\n", err)
			break
		}
		if err == io.EOF {
			fmt.Printf("已经读完成\n")
			break
		}
		if n == 0 {
			fmt.Printf("收到0字节数据\n")
			break
		}
		got := string(readData)
		fmt.Printf("接收到%v字节数据：%v\n", n, got)
		sn, err := conn.Write([]byte("服务器已经收到了"))
		if err != nil {
			fmt.Printf("发送失败：%v\n", err)
			break
		}
		fmt.Printf("发送成功%v字节\n", sn)

	}

}
