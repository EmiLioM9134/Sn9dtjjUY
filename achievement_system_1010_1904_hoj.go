// 代码生成时间: 2025-10-10 19:04:34
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

// Achievement 表示一个成就
type Achievement struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Points   int    `json:"points"`
    Unlocked bool   `json:"unlocked"`
}

// achievementStore 用于存储成就
var achievementStore = []Achievement{
    {ID: "ach1", Name: "First Login", Points: 10, Unlocked: false},
    {ID: "ach2", Name: "Complete Tutorial", Points: 20, Unlocked: false},
    {ID: "ach3", Name: "First Win", Points: 30, Unlocked: false},
}

// UnlockAchievement 处理解锁成就请求
func UnlockAchievement(c *gin.Context) {
    var req Achievement
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("Invalid request: %v", err),
        })
        return
    }
    for _, a := range achievementStore {
        if a.ID == req.ID {
            a.Unlocked = true
            c.JSON(http.StatusOK, gin.H{
                "message": "Achievement unlocked",
                "achievement": a,
            })
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{
        "error": "Achievement not found",
    })
}

// ListAchievements 列出所有成就
func ListAchievements(c *gin.Context) {
    c.JSON(http.StatusOK, achievementStore)
}

func main() {
    r := gin.Default()
    // 注册中间件
    // r.Use(gin.Recovery()) // 可选用的错误恢复中间件
    // 使用UnlockAchievement和ListAchievements处理器
    r.POST("/unlock", UnlockAchievement)
    r.GET("/achievements", ListAchievements)
    // 启动服务器
    r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
