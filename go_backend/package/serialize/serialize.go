package serialize

import (
	"fmt"
    "encoding/json"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    // "github.com/mitchellh/mapstructure"

    "strconv"
    "net/http"
   
)
type User struct {
    ID   uuid.UUID `gorm:"type:UUID;primaryKey"`
    Name string `gorm:"column:name;not null"`
    Age  int64    `gorm:"column:age;not null"`
}
type Server_online struct {
    Time   time.Time `gorm:"type:datetime;not null"`
    Data string `gorm:"column:data;not null"`
    // Age  int64    `gorm:"column:age;not null"`
}

// type Ve_monitor struct {
//     Time             time.Time `gorm:"type:datetime;not null"`
//     DA_status        string    `gorm:"column:da_status;not null"`
//     Data_reserved    string    `gorm:"column:data_reserved;not null"`
//     IPC_status       string    `gorm:"column:ipc_status"`
//     SH_status        string    `gorm:"column:sh_status"`
//     SW_version       string    `gorm:"column:sw_version"`
//     Vehicletype      string    `gorm:"column:vehicletype"`
//     Cpe_status       bool      `gorm:"column:cpe_status"`
//     Isp_lock         string    `gorm:"column:isp_lock"`
//     Isp_vsync        string    `gorm:"column:isp_vsync"`
//     Lidar_check      string    `gorm:"column:lidar_check"`
//     Monitor_pcie     string    `gorm:"column:monitor_pcie"`
//     Net_interface    string    `gorm:"column:net_interface"`
//     Net_status       string    `gorm:"column:net_status"`
//     Trunk_temperature float64      `gorm:"column:trunk_temperature"`
//     Vechicles_status string    `gorm:"column:vechicles_status"`
// }
type Ve_monitor struct {
    Time             time.Time `gorm:"type:datetime;not null"`
    DA_status        string    `gorm:"column:da_status;not null"`
    Data_reserved    string    `gorm:"column:data_reserved;not null"`
    IPC_status       string    `gorm:"column:ipc_status"`
	
	Plate_number     string    `gorm:"column:plate_number"`
    Source           string    `gorm:"column:source"`

    Type             string     `gorm:"column:type"`
	
    Schema_version   string      `gorm:"column:schema_version"`

    Vehicle          string     `gorm:"column:vehicle"`
    Fdi_type         string      `gorm:"column:fdi_type"`
    Timestamp        time.Time       `gorm:"column:timestamp"`

    SH_status        string    `gorm:"column:sh_status"`
    SW_version       string    `gorm:"column:sw_version"`
    Vehicletype      string    `gorm:"column:vehicletype"`
    Cpe_status       bool      `gorm:"column:cpe_status"`
    Isp_lock         string    `gorm:"column:isp_lock"`
    Isp_vsync        string    `gorm:"column:isp_vsync"`
    Lidar_check      string    `gorm:"column:lidar_check"`
    Monitor_pcie     string    `gorm:"column:monitor_pcie"`
    Net_interface    string    `gorm:"column:net_interface"`
    Net_status       string    `gorm:"column:net_status"`
    Trunk_temperature float64      `gorm:"column:trunk_temperature"`
    Vechicles_status string    `gorm:"column:vechicles_status"`
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
// func Insert_ve_monitor_data(c *gin.Context) {
//     var jsonMap map[string]interface{}

//     // Bind the JSON data from the request body
//     if err := c.ShouldBindJSON(&jsonMap); err != nil {
//         c.JSON(200, gin.H{"error": err.Error()})
//         return
//     }
    
//     // Create a new Ve_monitor instance
//     var veMonitor Ve_monitor
//     veMonitor.Time = time.Now()

//     // Loop through each item in the map and assign it to the corresponding field in Ve_monitor


    
   
//     c.JSON(200, gin.H{"your_post_data":"success"})
// }
// func Insert_ve_monitor_data(c *gin.Context){
//     var jsonMap map[string]interface{}

//     // 将请求体中的 JSON 绑定到 jsonMap 变量中
//     tt:=http.StatusBadRequest
//     fmt.Println(tt)
//     if err := c.ShouldBindJSON(&jsonMap); err != nil {
//         c.JSON(200, gin.H{"error": err.Error()})
//         return
//     }
//     jsonString, err := json.Marshal(jsonMap)
//     if err != nil {
//         fmt.Println("Error converting map to JSON:", err)
//         return
//     }

 

//     // 打印整个 JSON 对象
    
//     fmt.Println(string(jsonString))
//    // 遍历 map 提取所有的 key-value 对


//     ve_monitor:=??
//     result := Db.Create(&ve_monitor)
//     fmt.Println(result.RowsAffected)



   

// 	c.JSON(200, gin.H{"your_post_data":"success"})

// }
func Insert_ve_monitor_data(c *gin.Context){
    var jsonMap map[string]interface{}

    // 将请求体中的 JSON 绑定到 jsonMap 变量中
    tt:=http.StatusBadRequest
    fmt.Println(tt)
    if err := c.ShouldBindJSON(&jsonMap); err != nil {
        c.JSON(200, gin.H{"error": err.Error()})
        return
    }

    ve_monitor := Ve_monitor{}
    ve_monitor.Time = time.Now()  // 在这里设置 Time 字段为当前时间

    for key, value := range jsonMap {
        fmt.Println("kkkk:", key)
        fmt.Printf("Type of key %s is %T\n", key, value)

        switch v := value.(type) {
        case string:
            // 处理字符串类型的值
            switch key {
            case "type":
                ve_monitor.Type = v
            
            case "schema_version":
                ve_monitor.Schema_version = v
            case "SW_version":
                ve_monitor.SW_version = v
            case "Vehicletype":
                ve_monitor.Vehicletype = v
            
            case "vehicle":
                ve_monitor.Vehicle = v
            case "fdi_type":
                ve_monitor.Fdi_type = v
            case "plate_number":
                ve_monitor.Plate_number = v
            // case "Plate_number":
            //     ve_monitor.Plate_number = v
            case "source":
                ve_monitor.Source = v


        
            }
        case bool:
            if key == "Cpe_status" {
                ve_monitor.Cpe_status = v
            }
        case int64:
            {
                fmt.Println("timestamp----------------------------",v)
                if key == "timestamp" {
                    // ve_monitor.Timestamp = int64(v)
                    fmt.Println("timestamp----------------------------",v)
                }

            }
        case float64:
            if key == "trunk_temperature" {
                ve_monitor.Trunk_temperature = float64(v)
                fmt.Println("温度+++++++++++++++++++++++++++++++++++++",float64(v))
            }
            if key == "timestamp" {
                ve_monitor.Timestamp = time.Unix(int64(v), 0)
                // fmt.Println("温度+++++++++++++++++++++++++++++++++++++",float64(v))
            }
        case map[string]interface{}:
            // 处理嵌套的对象（如果有）
            jsonString, _ := json.Marshal(v)
            switch key {
            
            case "data_reserved":
                ve_monitor.Data_reserved = string(jsonString)
            case "DA_status":
                ve_monitor.DA_status = string(jsonString)
            case "IPC_status":
                ve_monitor.IPC_status = string(jsonString)
            case "SH_status":
                ve_monitor.SH_status = string(jsonString)
            case "isp_lock":
                ve_monitor.Isp_lock = string(jsonString)

            case "isp_vsync":
                ve_monitor.Isp_vsync = string(jsonString)
            case "lidar_check":
                ve_monitor.Lidar_check = string(jsonString)
             
            case "monitor_pcie":
                ve_monitor.Monitor_pcie = string(jsonString)
            case "net_interface":
                ve_monitor.Net_interface = string(jsonString)
            case "net_status":
                ve_monitor.Net_status = string(jsonString)
                fmt.Println("DA_status:", ve_monitor.Net_status)
            case "vechicles_status":
                ve_monitor.Vechicles_status = string(jsonString)
            
            
            }
            
        default:
            fmt.Println("Unexpected type found for key:", key)
        }
    }

    result := Db.Create(&ve_monitor)
    fmt.Println(result.RowsAffected)

    c.JSON(200, gin.H{"your_post_data":"success"})
}


func Insert_server_online(c *gin.Context) {
    var jsonMap map[string]interface{}

    // 将请求体中的 JSON 绑定到 jsonMap 变量中
    tt:=http.StatusBadRequest
    fmt.Println(tt)
    if err := c.ShouldBindJSON(&jsonMap); err != nil {
        c.JSON(200, gin.H{"error": err.Error()})
        return
    }
    jsonString, err := json.Marshal(jsonMap)
    if err != nil {
        fmt.Println("Error converting map to JSON:", err)
        return
    }

    // 打印整个 JSON 对象
    fmt.Println(jsonMap)
    fmt.Println(string(jsonString))
    // name:=c.Param("name")
    // data:=c.Param("data")
    
    // age,_:=strconv.ParseInt(c.Param("age"), 10, 64)

    // // 创建一个用户
    // user := User{ID:uuid.New(), Name: name, Age: age}
    // result := Db.Create(&user)
    // fmt.Println(result.RowsAffected)
    // fmt.Println(data)
    fmt.Printf("Type of data: %T\n", string(jsonString))

    server_online:=Server_online{Time:time.Now(), Data:string(jsonString)}
    result := Db.Create(&server_online)
    fmt.Println(result.RowsAffected)
   
    fmt.Println(c)



	c.JSON(200, gin.H{"your_post_data":"success"})
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
