package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
	})
}

func RedirectToAdd(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/add")
}

func GetSliceAny(c *gin.Context) {
	data := []any{"string", 123, true, map[string]any{"key": "value"}}
	c.JSON(http.StatusOK, data)
}

func GetSliceStruct(c *gin.Context) {
	data := []Item{
		{"item1", 1},
		{"item2", 2},
	}
	c.JSON(http.StatusOK, data)
}

func GetNumber(c *gin.Context) {
	c.JSON(http.StatusOK, 12345)
}

func GetString(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World!")
}

func GetMap(c *gin.Context) {
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	c.JSON(http.StatusOK, data)
}
