package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义模型
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);uniqueIndex"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"`
	IgnoreMe     int     `gorm:"-"`
}

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
	db.AutoMigrate(&UserInfo{}, &User{})

	flag := "User"
	if flag == "UserInfo" { // 创建数据行
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
	} else {
		// u := &User{
		// 	Name: "kimi",
		// }
		// // 创建
		// db.Create(&u)
		// var u1 User

		// // 查询
		// db.Create(&u1)

	}
}
