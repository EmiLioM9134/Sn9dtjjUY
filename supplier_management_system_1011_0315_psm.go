// 代码生成时间: 2025-10-11 03:15:24
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Supplier 代表供应商的结构体
type Supplier struct {
    ID         int    `json:"id"`
    Name       string `json:"name"`
    Contact    string `json:"contact"`
    PhoneNumber string `json:"phone_number"`
}

// GetAllSuppliers 处理获取所有供应商的请求
func GetAllSuppliers(c *gin.Context) {
    suppliers := []Supplier{
        {ID: 1, Name: "Supplier A", Contact: "John Doe", PhoneNumber: "1234567890"},
        {ID: 2, Name: "Supplier B", Contact: "Jane Doe", PhoneNumber: "0987654321"},
    }
    // 如果供应商列表为空，返回错误
    if len(suppliers) == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "No suppliers found",
        })
        return
    }
    c.JSON(http.StatusOK, suppliers)
}

// CreateSupplier 处理创建新供应商的请求
func CreateSupplier(c *gin.Context) {
    var newSupplier Supplier
    if err := c.ShouldBindJSON(&newSupplier); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    // 这里可以添加逻辑来保存供应商信息到数据库
    // 模拟保存新供应商
    c.JSON(http.StatusCreated, gin.H{
        "message": "Supplier created successfully",
        "supplier": newSupplier,
    })
}

func main() {
    r := gin.Default()

    // 使用中间件记录请求日志
    r.Use(gin.Logger())

    // 使用中间件恢复任何panic，将程序从崩溃中恢复
    r.Use(gin.Recovery())

    // 定义路由
    r.GET("/suppliers", GetAllSuppliers)
    r.POST("/suppliers", CreateSupplier)

    // 启动服务器
    r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
