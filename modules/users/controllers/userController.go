package controllers

import "github.com/gin-gonic/gin"

func HelloUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello User",
	})
}
