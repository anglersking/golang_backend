package main
import (
    "fmt"
    // "gorm.io/driver/clickhouse"
    // "gorm.io/gorm"
    "github.com/gin-gonic/gin"
    "go_project/package/serialize"
)


func main() {
    fmt.Println("start gorm")
    // 连接 ClickHouse 数据库
    router := gin.Default()
    router.POST("/user", serialize.CreateUser)
	router.GET("/user/:id", serialize.GetUser)
    router.GET("/user", serialize.CreateUser)

	router.Run(":8089")
   
    


 
}
// package main

// import (
// 	"fmt"
// 	"github.com/gin-gonic/gin"
// 	"go_project/package/serialize"
// 	"sync"
// )

// func main() {
// 	fmt.Println("start gorm")
// 	// 连接 ClickHouse 数据库
// 	router := gin.Default()

// 	var wg sync.WaitGroup // 定义一个WaitGroup来等待各个Goroutines完成
// 	wg.Add(3)             // 添加3个Goroutines

// 	userCreateCh := make(chan bool)
// 	userGetCh := make(chan bool)
// 	userListCh := make(chan bool)

// 	go func() {
// 		router.POST("/user", serialize.CreateUser)
// 		userCreateCh <- true
// 		wg.Done()
// 	}()

// 	go func() {
// 		router.GET("/user/:id", serialize.GetUser)
// 		userGetCh <- true
// 		wg.Done()
// 	}()

// 	go func() {
// 		router.GET("/user", serialize.CreateUser)
// 		userListCh <- true
// 		wg.Done()
// 	}()

// 	go router.Run(":8089")

// 	<-userCreateCh
// 	<-userGetCh
// 	<-userListCh

// 	wg.Wait() // 等待所有Goroutines完成
// }
