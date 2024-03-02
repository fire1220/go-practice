package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1888")
	if err != nil {
		log.Printf("建立链接失败,err:%v\n", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("关闭链接失败，err:%v\n", err)
		}
	}()

	sn, err := conn.Write([]byte("hello world"))
	if err != nil {
		log.Printf("发送数据失败，err:%v\n", err)
		return
	}
	log.Printf("发送%v个字节\n", sn)

	readData := make([]byte, 1024)
	n, err := conn.Read(readData)
	if err != nil {
		log.Printf("读取数据失败,err:%v\n", err)
		return
	}
	log.Printf("读取%v个字节，内容：%v\n", n, string(readData))
}
