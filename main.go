package main

import (
	"QingWork/model"
	"fmt"
	"gorm.io/gen"
)

func main() {
	err := model.DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("自动迁移 User 表失败！")
		return
	}
	fmt.Println("自动迁移 User 表成功！")

	// 自动生成 struct
	GeneratePeopleStruct()
}

// GeneratePeopleStruct 生成 People 结构体
func GeneratePeopleStruct() {
	g := gen.NewGenerator(gen.Config{
		OutFile: "/tmp/QingWord/query",
	})

	g.UseDB(model.DB)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
