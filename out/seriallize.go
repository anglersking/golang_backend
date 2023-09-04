package serialize




func createUser(c *gin.Context) {

	// var newUser User
	// if err := c.ShouldBindJSON(&newUser); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	// users = append(users, newUser)
    

    dsn := "clickhouse+native://default:7ec01302-c218-11ed-bfe5-07b35aad18b4@122.9.166.18:9000/center_ods"
    db, err := gorm.Open(clickhouse.New(clickhouse.Config{
        DSN: dsn,
    }), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 自动迁移数据表，如果不存在则创建
    db.AutoMigrate(&User{})

    // 创建一个用户
    user := User{Name: "Alice", Age: 20}
    result := db.Create(&user)
    fmt.Println(result.RowsAffected)

    // 查询所有用户
    var users []User
    db.Find(&users)
    fmt.Println(users)

    // 查询单个用户
    var u User
    db.First(&u, "name = ?", "Alice")
    fmt.Println(u)

    // 更新用户信息
    db.Model(&user).Where("id = ?", user.ID).Update("Age", 999)

	c.JSON(200, gin.H{"status": "User created", "user": "heihei"})
}
func getUser(c *gin.Context) {
	id := c.Param("id")

	// for _, user := range users {
	// 	if id == strconv.FormatUint(user.ID, 10) {
	// 		c.JSON(200, gin.H{"user": user})
	// 		return
	// 	}
	// }

    dsn := "clickhouse+native://default:7ec01302-c218-11ed-bfe5-07b35aad18b4@122.9.166.18:9000/center_ods"
    db, err := gorm.Open(clickhouse.New(clickhouse.Config{
        DSN: dsn,
    }), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }


    var u User
    db.First(&u, "name = ?", "Alice")
    fmt.Println(u)
       // 删除用户
    db.Where("id = ?", id).Delete(&u)

	c.JSON(200, gin.H{"status": id})
}