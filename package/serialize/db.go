package serialize

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

var Db *gorm.DB

func SetupDatabase() {
	dsn := "host=106.15.77.79 user=postgres password=123456 dbname=test port=5432"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	fmt.Println("Database connection established successfully")

	// 自动迁移数据表，如果不存在则创建
	Db.AutoMigrate(&Task_info{})
}
