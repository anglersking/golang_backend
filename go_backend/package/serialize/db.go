package serialize

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"fmt"
)

var Db *gorm.DB

func SetupDatabase() {
	dsn := "clickhouse+native://default:7ec01302-c218-11ed-bfe5-07b35aad18b4@10.9.61.34:9002/server_test_ods"
	var err error
	Db, err = gorm.Open(clickhouse.New(clickhouse.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
   fmt.Println("create table success")

	// 自动迁移数据表，如果不存在则创建
	Db.AutoMigrate(&Server_online{})
}