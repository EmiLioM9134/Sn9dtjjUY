// 代码生成时间: 2025-09-24 12:28:36
package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MathTool 结构体封装数学计算功能
type MathTool struct {}

// Add 两个数相加
func (t *MathTool) Add(c *gin.Context) {
	first, err := strconv.ParseFloat(c.PostForm("first"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid first argument",
		})
		return
	}
	second, err := strconv.ParseFloat(c.PostForm("second"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid second argument",
		})
		return
	}
	result := first + second
	c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
}

// Subtract 两个数相减
func (t *MathTool) Subtract(c *gin.Context) {
	first, err := strconv.ParseFloat(c.PostForm("first"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid first argument",
		})
		return
	}
	second, err := strconv.ParseFloat(c.PostForm("second"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid second argument",
		})
		return
	}
	result := first - second
	c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
}

// Multiply 两个数相乘
func (t *MathTool) Multiply(c *gin.Context) {
	first, err := strconv.ParseFloat(c.PostForm("first"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid first argument",
		})
		return
	}
	second, err := strconv.ParseFloat(c.PostForm("second"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid second argument",
		})
		return
	}
	result := first * second
	c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
}

// Divide 两个数相除
func (t *MathTool) Divide(c *gin.Context) {
	first, err := strconv.ParseFloat(c.PostForm("first"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid first argument",
		})
		return
	}
	second, err := strconv.ParseFloat(c.PostForm("second"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid second argument",
		})
		return
	}
	if second == 0.0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Division by zero",
		})
		return
	}
	result := first / second
	c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
}

func main() {
	r := gin.Default()
	r.POST("/add", (&MathTool{}).Add)
	r.POST("/subtract", (&MathTool{}).Subtract)
	r.POST("/multiply", (&MathTool{}).Multiply)
	r.POST("/divide", (&MathTool{}).Divide)
	r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
