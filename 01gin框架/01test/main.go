package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//示例1：
	//r := gin.Default()
	//r.POST("/",func(c *gin.Context){
	//	c.String(http.StatusOK,"hello world!")
	//})
	//r.Run(":8000")

	//示例2
	//r := gin.Default()
	//r.GET("/",func(c *gin.Context){
	//	c.String(http.StatusOK,"hello world!")
	//})
	//r.POST("/xxxPost",getting)
	//r.PUT("/xxxPut")
	//r.Run(":8080")

	//示例3
	//r := gin.Default()
	//r.GET("/user/:name/*action",func(c *gin.Context){
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	c.String(http.StatusOK,name+" is "+action)
	//})
	//r.Run(":8080")

	//示例4
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context){
		name := c.DefaultQuery("sb","Jack")
		c.String(http.StatusOK,fmt.Sprintf("hello %s",name))
	})
	r.Run(":8080")

}
