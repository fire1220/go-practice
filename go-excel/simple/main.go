package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"unsafe"
)

func main() {
	openFile, err := excelize.OpenFile("file.xlsx") // 读取文件
	if err != nil {
		log.Printf("OpenFile err:%v\n", err)
		return
	}
	excelRowList, err := openFile.GetRows("Sheet1") // 读取行
	if err != nil {
		log.Printf("GetRows err:%v\n", err)
		return
	}
	j, err := json.Marshal(excelRowList)
	if err != nil {
		log.Printf("Marshal err:%v\n", err)
		return
	}
	fmt.Println(*(*string)(unsafe.Pointer(&j)))
}
