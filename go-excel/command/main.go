package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"unsafe"
)

var fileName string  // 文件名称
var sheetName string // 表格的sheet名称

func main() {
	flag.StringVar(&fileName, "f", "", "表格文件路径，例如 file.xlsx")
	flag.StringVar(&sheetName, "s", "Sheet1", "表的sheet名称，默认是Sheet1")
	flag.Parse() // 解析命令行
	if fileName == "" {
		fmt.Println("没有对应文件，参数是 -f")
		return
	}
	openFile, err := excelize.OpenFile(fileName) // 读取文件
	if err != nil {
		log.Printf("OpenFile err:%v\n", err)
		return
	}
	excelRowList, err := openFile.GetRows(sheetName) // 读取行
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
