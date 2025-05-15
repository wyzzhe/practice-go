package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func Init() {
	// gorm-mysql
	dsn := "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("fail to connect database: %v", err)
	}

	// 自动迁移表结构
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	u1 := &UserInfo{
		ID:     1,
		Name:   "kimi",
		Gender: "男",
		Hobby:  "双色球",
	}
	db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u)

	// 更新
	db.Model(&u).Update("hobby", "篮球")

	db.First(&u)
	fmt.Printf("%#v \n", u)

	// 删除
	db.Delete(&u)
}
