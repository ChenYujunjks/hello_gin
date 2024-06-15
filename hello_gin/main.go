package main

import (
	"fmt"
	routes "hello_gin/routers"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdditionRequest struct {
	Number1 float64 `json:"number1" binding:"required"`
	Number2 float64 `json:"number2" binding:"required"`
}

func main() {
	gin.SetMode(gin.DebugMode) // 设置Gin的运行模式为DebugMode

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	routes.RegisterRoutes(r)

	//查询参数绑定
	r.GET("/multiply", func(c *gin.Context) {
		number1Str := c.Query("number1")
		number2Str := c.Query("number2")

		// 如果没有提供查询参数，则渲染HTML表单
		if number1Str == "" && number2Str == "" {
			c.HTML(http.StatusOK, "multiply.html", nil)
			return
		}
		// 处理提交的查询参数
		if number1Str == "" || number2Str == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "number1 and number2 are required"})
			return
		}
		number1, err1 := strconv.ParseFloat(number1Str, 64)
		number2, err2 := strconv.ParseFloat(number2Str, 64)

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number format"})
			return
		}

		result := number1 * number2
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add.html", nil)
	})
	r.GET("/subtract", func(c *gin.Context) {
		c.HTML(http.StatusOK, "subtract.html", nil)
	})
	//form 参数绑定
	r.POST("/subtract", func(c *gin.Context) {
		number1 := c.PostForm("number1")
		number2 := c.PostForm("number2")

		if number1 == "" || number2 == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "number1 and number2 are required"})
			return
		}
		num1, err1 := strconv.ParseFloat(number1, 64)
		num2, err2 := strconv.ParseFloat(number2, 64)
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number format"})
			return
		}
		result := num1 - num2
		c.JSON(http.StatusOK, gin.H{"result": result})
	})
	//JSON 参数示例
	r.POST("/add", func(c *gin.Context) {
		number1 := c.PostForm("number1")
		number2 := c.PostForm("number2")

		if number1 == "" || number2 == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "number1 and number2 are required"})
			return
		}
		num1, err1 := strconv.ParseFloat(number1, 64)
		num2, err2 := strconv.ParseFloat(number2, 64)
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number format"})
			return
		}
		result := num1 + num2
		c.JSON(http.StatusOK, gin.H{"result": result})
	})
	// 处理GET请求，渲染HTML页面
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})
	r.Run(":8080")
}
