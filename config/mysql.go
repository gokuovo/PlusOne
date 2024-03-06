package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

func mysqlCon(mysqlDetail DatabaseConfig) {

	// 连接数据库
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlDetail.Username, mysqlDetail.Password, mysqlDetail.Host, mysqlDetail.Port, mysqlDetail.Dbname)
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
	//err = DB.AutoMigrate(&User{})
	//if err != nil {
	//	log.Fatal("Database migration failed:", err)
	//}
}
