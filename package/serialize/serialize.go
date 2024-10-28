package serialize

import (
    "encoding/json"
    // "fmt"
    "net/http"
    "strconv"
    // "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    ID   uuid.UUID `gorm:"type:UUID;primaryKey"`
    Name string    `gorm:"column:name;not null"`
    Age  int64     `gorm:"column:age;not null"`
}

type TaskInfoResponse struct {
    Time      int64   `json:"time"`
    PhoneNumber string      `json:"phonenumber"`
    JobStatus int64       `json:"jobstatus"`
    JobInfo   interface{} `json:"jobinfo"` // 使用 interface{} 以支持解析后的 JSON
}

type Task_info struct {
    Time     int64  `gorm:"column:time;not null;primaryKey"` // 毫秒级时间戳
    PhoneNumber string `gorm:"column:phonenumber;not null"`
    JobStatus int64 `gorm:"column:jobstatus;not null"`
    JobInfo  string `gorm:"column:jobinfo;not null"`
}

// 创建任务
func Create_Task(c *gin.Context) {
    var task Task_info
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 假设传入的时间戳是毫秒级别的int64
    result := Db.Create(&task)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "task": task})
}

// 获取任务信息
func GetTaskInfo(c *gin.Context) {
    var tasks []Task_info
    if err := Db.Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var response []TaskInfoResponse
    for _, task := range tasks {
        var jobInfo interface{}
        // 解析 JobInfo 字段为 JSON
        if err := json.Unmarshal([]byte(task.JobInfo), &jobInfo); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JobInfo"})
            return
        }

        // 将 int64 时间戳转换为 time.Time
        // taskTime := time.Unix(task.Time/1000, (task.Time%1000)*1000000) // 毫秒转为秒和纳秒

        response = append(response, TaskInfoResponse{
            Time:      task.Time,
            PhoneNumber:  task.PhoneNumber,
            JobStatus: task.JobStatus,
            JobInfo:   jobInfo,
        })
    }

    c.JSON(http.StatusOK, response)
}

// 根据时间戳和用户名获取任务
func GetTaskByTimeAndUser(c *gin.Context) {
    timeParam := c.Query("time")   // 从查询参数获取毫秒级时间戳字符串
    phonenumber := c.Query("phonenumber") // 从查询参数获取用户名

    // 将时间戳字符串转换为 int64
    timestamp, err := strconv.ParseInt(timeParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
        return
    }

    var task Task_info
    if err := Db.Where("time = ? AND phonenumber = ?", timestamp, phonenumber).First(&task).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    var jobInfo interface{}
    // 解析 JobInfo 字段为 JSON
    if err := json.Unmarshal([]byte(task.JobInfo), &jobInfo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JobInfo"})
        return
    }

    // 将 int64 时间戳转换为 time.Time
    // taskTime := time.Unix(task.Time/1000, (task.Time%1000)*1000000)

    response := TaskInfoResponse{
        Time:      task.Time,
        PhoneNumber:  task.PhoneNumber,
        JobStatus: task.JobStatus,
        JobInfo:   jobInfo,
    }

    c.JSON(http.StatusOK, response)
}
