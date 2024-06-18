package controllers

import (
	"login_mvc/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var user models.User
	var requestUser models.User

	// 解析请求体中的JSON数据
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库中的用户信息
	if err := ctrl.DB.Where("username = ? AND password = ?", requestUser.Username, requestUser.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "欢迎, " + user.Username})
}
