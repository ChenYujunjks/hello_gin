package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdditionRequest struct {
	Number1 float64 `json:"number1" binding:"required"`
	Number2 float64 `json:"number2" binding:"required"`
}

func getStudentData() gin.H {
	return gin.H{
		"name": "王石头",
	}
}
func studentHandler(c *gin.Context) {
	data := getStudentData()
	c.JSON(http.StatusOK, data)
}
func schoolHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"school": "idk", "students": getStudentData()})
}
func main() {
	gin.SetMode(gin.DebugMode) // 设置Gin的运行模式为DebugMode

	r := gin.Default()
	r.LoadHTMLGlob("../templates/*")

	r.GET("/student", studentHandler)
	r.GET("/school", schoolHandler)

	// JSON 用法
	r.GET("/ping3", func(c *gin.Context) {
		type Response struct {
			Message string `json:"message"`
		}
		c.JSON(http.StatusOK, Response{
			Message: "pong",
		})
	})
	r.Run(":8080")
}
