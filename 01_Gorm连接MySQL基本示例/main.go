package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	username := "root"
	password := "yjb@1234"
	host := "127.0.0.1"
	port := 3306
	Dbname := "db1"  // 数据库名
	timeout := "10s" // 连接超时，10s

	// root:root@tcp(127.0.0.1:3306)/gorm?
	// 设置字符集charset（Mysql推荐使用）,设置解析时间类型字段（如DATA,DATETIME,TIMESTAMP）的方法，为True则将其自动解析为Go的time.Time类型。loc设置数据库连接的时区为本地时区
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	// 连接Mysql，获取DB类型实例，用于后续的数据库读写操作
	// 传递给Open的第二个参数，表示Gorm的一些配置
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败，error=" + err.Error())
	}

	// Gorm v2中不再显式提供关闭数据库连接的操作，Gorm会在程序结束时自动释放数据库连接资源
	// defer db.Close()

	// 创建表结构 自动迁移（将结构体与数据表进行对应）
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		panic(err)
	}

	// 创建数据行
	/*
		u1 := UserInfo{
				ID:     1,
				Name:   "七米",
				Gender: "男",
				Hobby:  "蛙泳",
			}
		db.Create(&u1)
	*/

	// 查询表中第一条数据保存到u中
	var u UserInfo
	db.First(&u)
	fmt.Printf("%#v\n", u)

	// 更新
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)
}
