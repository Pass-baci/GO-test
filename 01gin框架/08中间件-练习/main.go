package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Next()
		since := time.Since(now)
		fmt.Println("time:",since)
	}
}

func main() {
	r := gin.Default()
	v1 := r.Group("v1")
	v1.Use(MiddleWare())
	{
		v1.GET("/index", index)
		v1.GET("/index1", index1)
	}
	r.Run(":8080")
}

func index(c *gin.Context) {
	for i := 0; i<100000000; i++ {}
	c.String(http.StatusOK,"nihao")
}

func index1(c *gin.Context) {
	for i := 0; i<1000000001; i++ {}
	c.String(http.StatusOK,"nihao")
}
