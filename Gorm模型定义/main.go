package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// User 定义模型
type User struct {
	gorm.Model   // 内嵌
	Name         string
	Age          sql.NullInt64 `gorm:"column:age"` // 零值类型（此Tag可以给列改名字，对于创建新表则是直接名字替换，对于旧表的操作则是另创建一列名字叫这个，之前那个还是保留的）
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置num为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

func (User) TableName() string {
	return "User"
}

// Animal 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// TableName 唯一指定表名优先级最高（自定义表明）
func (Animal) TableName() string {
	return "animals"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//NamingStrategy: schema.NamingStrategy{ // 禁用表名复数
		//	SingularTable: true,
		//},
	})
	if err != nil {
		panic("连接数据库失败，error=" + err.Error())
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	/*
		err = db.AutoMigrate(&Animal{})
		if err != nil {
			panic(err)
		}
	*/

}
