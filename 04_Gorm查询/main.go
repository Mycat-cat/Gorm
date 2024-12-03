package main

import (
	"Gorm/DB"
	"time"
)

func main() {
	db := DB.ConnectDB()
	/*
		err := db.AutoMigrate(&DB.User{})
		if err != nil {
			panic(err)
		}
	*/

	user := DB.User{
		Name:     "firstRecord",
		Age:      18,
		Birthday: time.Now(),
	}

	// 创建记录并为指定字段赋值
	//db.Select("Name", "Age", "CreatedAt").Create(&user)

	//创建记录并忽略传递给'Omit'的字段值
	db.Omit("Name", "Age", "CreatedAt", "UpdateAt").Create(&user)
}
