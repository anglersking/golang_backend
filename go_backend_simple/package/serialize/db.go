package serialize

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var Db *gorm.DB
type User struct {
    ID   uuid.UUID `gorm:"type:UUID;primaryKey"`
    Name string `gorm:"column:name;not null"`
    Age  int64    `gorm:"column:age;not null"`
}

func SetupDatabase() {
	fmt.Println("creat success")
	dsn := "clickhouse+native://default:7ec01302-c218-11ed-bfe5-07b35aad18b4@122.9.166.18:9000/center_ods"
	var err error
	Db, err = gorm.Open(clickhouse.New(clickhouse.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移数据表，如果不存在则创建
	Db.AutoMigrate(&User{})
	fmt.Println("creat success")
}