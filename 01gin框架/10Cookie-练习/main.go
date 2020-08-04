package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("key"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error":"err"})
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()
	{
		r.GET("/login", func(c *gin.Context) {
			c.SetCookie("key", "123", 60, "/", "localhost", false, true)
			c.String(http.StatusOK, "login success")
		})
		r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
			c.JSON(200,gin.H{"data":"home"})
		})
	}
	r.Run(":8080")
}
