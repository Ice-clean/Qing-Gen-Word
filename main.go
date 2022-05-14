package main

import (
	"QingWork/model"
	_ "QingWork/model"
	"fmt"
)

func main() {
	err := model.DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("自动迁移 User 表失败！")
	}
}
