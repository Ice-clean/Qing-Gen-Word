package model

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
)

// 数据库配置
type mysqlConfig struct {
	Addr   string `yaml:"addr"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	DBName string `yaml:"db-name"`
	DSN    string
}

var DB *gorm.DB

func init() {
	config := &mysqlConfig{}
	config.loadConfigFile("/tmp/config.yml")
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("打开数据库失败", err)
	}
	DB = db
}

// 读取 yaml 配置文件
func (m *mysqlConfig) loadConfigFile(path string) {
	// 读取 yaml 文件
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("打开 yaml 文件失败", err)
		return
	}

	// 解析 yaml 文件
	err = yaml.Unmarshal(file, m)
	if err != nil {
		fmt.Println("解析 yaml 文件失败", err)
	}

	// 拼接 dsn
	m.DSN = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Passwd, m.Addr, m.DBName)
}
