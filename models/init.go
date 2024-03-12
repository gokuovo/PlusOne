package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

func MysqlCon(dsn string) {
	// 连接数据库
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Database connected successfully!")
	// 自动迁移数据库表结构（创建表）
	err = DB.AutoMigrate(&Test{})
	if err != nil {
		log.Fatal("Database migration failed:", err)
	}
}
