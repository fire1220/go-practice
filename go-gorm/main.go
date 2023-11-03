package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/model"
)

func main() {
	dsn := `web_user:l%meFN!Z88yRgrjz@tcp(rm-2zeti0v9e6940n93p.mysql.rds.aliyuncs.com)/wm_light_backend?charset=utf8mb4&parseTime=true&loc=Local`
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	data := model.WmDataCorrection{
		Mid: 123, // 学员id
	}
	err = db.Create(&data).Error
	if err != nil {
		fmt.Println("数据执行失败", err)
	}
	fmt.Println(data)
}
