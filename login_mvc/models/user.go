package models

import "gorm.io/gorm"

// 这个文件不需要修改，因为数据模型和数据库类型无关。
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	//nigeli   string `form:"username" json:"password" binding:"required"`
}
