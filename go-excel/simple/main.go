package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"unsafe"
)

func main() {
	j, err := readToJson("file.xlsx")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(j)
	}
}

// 读取表格,返回json
func readToJson(fileName string, sheetL ...string) (string, error) {
	sheet := "Sheet1"
	if len(sheetL) > 0 {
		sheet = sheetL[0]
	}
	openFile, err := excelize.OpenFile(fileName) // 读取文件
	if err != nil {
		return "", fmt.Errorf("OpenFile err:%v\n", err)
	}
	excelRowList, err := openFile.GetRows(sheet) // 读取行
	if err != nil {
		return "", fmt.Errorf("GetRows err:%v\n", err)
	}
	j, err := json.Marshal(excelRowList)
	if err != nil {
		return "", fmt.Errorf("Marshal err:%v\n", err)
	}
	return *(*string)(unsafe.Pointer(&j)), nil
}

// 写入表格
func write() {

}
