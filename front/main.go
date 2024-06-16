package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 首先加载templates目录下面的所有模版文件，模版文件扩展名随意
	router.LoadHTMLGlob("templates/*")

	//  /assets/images/1.jpg 这个url文件，存储在/public/images/1.jpg
	router.Static("/assets", "public")

	// 为单个静态资源文件，绑定url
	// 这里的意思就是将/favicon.ico这个url，绑定到./resources/favicon.ico这个文件
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		// 通过HTML函数返回html代码
		// 第二个参数是模版文件名字
		// 第三个参数是map类型，代表模版参数
		// gin.H 是map[string]interface{}类型的别名
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Yujun Chen's Personal Website",
		})
	})
	router.Run(":8080")
}
