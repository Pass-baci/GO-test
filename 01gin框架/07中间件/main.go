package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200,"nihao")
		//c.Abort()
	}
}
func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200,"nihao1")
		//c.Abort()
	}
}

func main() {
	r := gin.Default()
	v1 := r.Group("v1")
	v2 := r.Group("v2")
	v2.Use(MiddleWare())
	{
		v1.GET("/middleware1",MiddleWare2(), func(c *gin.Context) {
			c.String(http.StatusOK,"hello")
		})
	}
	{
		v2.GET("/middleware", func(c *gin.Context) {
			c.String(http.StatusOK,"hello")
		})
	}
	//for i:= 0; i<100000000; i++ {}
	r.Run(":8080")
}
