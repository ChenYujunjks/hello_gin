package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 注册各个路由模块
	RegisterUserRoutes(r)
	RegisterAuthRoutes(r)
}
