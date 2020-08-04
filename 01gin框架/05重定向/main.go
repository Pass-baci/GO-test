package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/home", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})
	r.Run(":8080")
}
