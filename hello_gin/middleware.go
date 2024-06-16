package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Next 中间件都是栈 先进后出

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/add" {
			c.Next()
			return
		}
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println("------------", cost)
	}
}
