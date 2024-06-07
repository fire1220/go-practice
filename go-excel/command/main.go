package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"unsafe"
)

var fileName string  // 文件名称
var sheetName string // 表格的sheet名称

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "%s 将表格文件转换为json\n\n", flag.CommandLine.Name())
		fmt.Fprintf(os.Stdout, "Usage: %s example.xlsx\n", flag.CommandLine.Name())
		fmt.Fprintf(os.Stdout, "       %s -f example.xlsx\n", flag.CommandLine.Name())
		fmt.Fprintf(os.Stdout, "       %s -s sheet1 -f example.xlsx\n\n", flag.CommandLine.Name())
		flag.PrintDefaults()
	}
	flag.StringVar(&fileName, "f", "", "表格文件路径，例如 example.xlsx")
	flag.StringVar(&sheetName, "s", "", "表格sheet名称，默认是第一个sheet")
	flag.Parse() // 解析命令行

	if fileName == "" && len(flag.Args()) > 0 {
		fileName = flag.Args()[len(flag.Args())-1]
	}

	if fileName == "" {
		flag.Usage()
		return
	}

	openFile, err := excelize.OpenFile(fileName) // 读取文件
	if err != nil {
		fmt.Printf("OpenFile err:%v\n", err)
		return
	}

	defer func(openFile *excelize.File) {
		err := openFile.Close()
		if err != nil {
			fmt.Printf("OpenFile.Close err:%v\n", err)
		}
	}(openFile)

	if sheetName == "" {
		sheetName = openFile.GetSheetName(0)
	}

	excelRowList, err := openFile.GetRows(sheetName) // 读取行
	if err != nil {
		fmt.Printf("GetRows err:%v\n", err)
		return
	}

	// 遍历行数据，确保空单元格返回空字符串
	firstRowCount := 0
	for k, row := range excelRowList {
		if k == 0 {
			firstRowCount = len(row)
			continue
		}
		full := firstRowCount - len(row)
		if full > 0 {
			for i := 0; i < full; i++ {
				row = append(row, "")
			}
			excelRowList[k] = row
		}
	}

	j, err := json.Marshal(excelRowList)
	if err != nil {
		fmt.Printf("Marshal err:%v\n", err)
		return
	}
	fmt.Println(*(*string)(unsafe.Pointer(&j)))
}
