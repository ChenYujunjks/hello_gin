package controllers

import (
	"log"
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
		log.Println((err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库中的用户信息
	if err := ctrl.DB.Where("username = ? AND password = ?", requestUser.Username, requestUser.Password).First(&user).Error; err != nil {
		log.Println((err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "欢迎, " + user.Username})
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var requestUser models.User

	// 解析请求体中的JSON数据
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存用户信息到数据库
	if err := ctrl.DB.Create(&requestUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func (ctrl *AuthController) GetUsers(c *gin.Context) {
	var users []models.User
	if err := ctrl.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户列表"})
		return
	}
	c.HTML(http.StatusOK, "users.html", gin.H{"users": users})
}
func (ctrl *AuthController) ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (ctrl *AuthController) ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
