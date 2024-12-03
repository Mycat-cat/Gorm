package DB

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
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
	return db
}
