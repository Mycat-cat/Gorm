package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int
	Name sql.NullString `gorm:"default:'小王子'"`
	Age  int64          `gorm:"default:20"`
}

func main() {
	// 连接数据库
	username := "root"
	password := "yjb@1234"
	host := "127.0.0.1"
	port := 3306
	Dbname := "db1"  // 数据库名
	timeout := "10s" // 连接超时，10s
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败，error=" + err.Error())
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	u := User{
		Name: sql.NullString{
			String: "",
			Valid:  true,
		},
		Age: 188,
	}
	db.Debug().Create(&u)
	//db.Create(&u)
}
