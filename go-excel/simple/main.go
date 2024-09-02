package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"time"
	"unsafe"
)

func main() {
	tWriteAppendSave()
	tReadToJson()
}

func tReadToJson() {
	j, err := readToJson("file.xlsx")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(j)
	}
}

func tWriteAppendSave() {
	fileName := "file.xlsx"
	var data [][]interface{}
	data = append(data, []interface{}{1, "jock", 12})
	data = append(data, []interface{}{2, "fire", 13})
	e := writeAppendSave(fileName, data, "Sheet1")
	if e != nil {
		fmt.Println(fmt.Errorf("第1次批量写入失败：err:%v", e))
		return
	}
	data = [][]interface{}{}
	time.Sleep(10 * time.Second)
	data = append(data, []interface{}{3, "zhangSan", 14})
	data = append(data, []interface{}{4, "lisi", 15})
	e = writeAppendSave(fileName, data, "Sheet1", 3)
	if e != nil {
		fmt.Println(fmt.Errorf("第2次批量写入失败：err:%v", e))
		return
	}
	return
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
func writeAppendSave(fileName string, data [][]interface{}, sheet string, appendRowNumL ...int) error {
	appendRowNum := 1
	if len(appendRowNumL) > 0 {
		appendRowNum = appendRowNumL[0]
	}
	var f *excelize.File
	var err error
	if fileExists(fileName) {
		f, err = excelize.OpenFile(fileName)
		if err != nil {
			return err
		}
	} else {
		f = excelize.NewFile()
		index, err := f.NewSheet(sheet)
		if err != nil {
			return fmt.Errorf("创建失败；err:%v", err)
		}
		f.SetActiveSheet(index)
	}
	for _, val := range data {
		err = f.SetSheetRow(sheet, fmt.Sprintf("A%d", appendRowNum), &val)
		if err != nil {
			return fmt.Errorf("写入失败；err:%v", err)
		}
		appendRowNum++
	}
	return f.SaveAs(fileName)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
