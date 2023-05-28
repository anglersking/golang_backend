package serialize

import (
	"fmt"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "strconv"
   
)
type User struct {
    ID   uuid.UUID `gorm:"type:UUID;primaryKey"`
    Name string `gorm:"column:name;not null"`
    Age  int64    `gorm:"column:age;not null"`
}

func Getall(c *gin.Context) {

    // 查询所有用户
    var users []User
    Db.Find(&users)
    fmt.Println(users)
    c.JSON(200, users)

    // // 查询单个用户
    // var u User
    // Db.First(&u, "name = ?", "Alice")
    // fmt.Println(u)

    
}

func CreateUser(c *gin.Context) {

    name:=c.Param("name")
    
    age,_:=strconv.ParseInt(c.Param("age"), 10, 64)

    // 创建一个用户
    user := User{ID:uuid.New(), Name: name, Age: age}
    result := Db.Create(&user)
    fmt.Println(result.RowsAffected)


	c.JSON(200, gin.H{"status": "User created success"})
}
func GetUser(c *gin.Context) {
	name := c.Param("name")



    var u User
    Db.First(&u, "name = ?", name)
    fmt.Println(u)
   

	c.JSON(200, u)
}

func DeleteUser(c *gin.Context) {
	name := c.Param("name")
    

	

    var u User
  
    Db.Where("name = ?", name).Delete(&u)

	c.JSON(200, gin.H{"status": "delete success"})
}
func UpdateUser(c *gin.Context) {
	name := c.Param("name")
    age,_:=strconv.ParseInt(c.Param("age"), 10, 64)



	// for _, user := range users {
	// 	if id == strconv.FormatUint(user.ID, 10) {
	// 		c.JSON(200, gin.H{"user": user})
	// 		return
	// 	}
	// }

    // dsn := "clickhouse+native://default:7ec01302-c218-11ed-bfe5-07b35aad18b4@122.9.166.18:9000/center_ods"
    // Db, err := gorm.Open(clickhouse.New(clickhouse.Config{
    //     DSN: dsn,
    // }), &gorm.Config{})
    // if err != nil {
    //     panic("failed to connect database")
    // }


    var user []User
  
   

    Db.Model(&user).Where("name = ?", name).Update("Age", age)

	c.JSON(200, gin.H{"status": "update success"})
}
