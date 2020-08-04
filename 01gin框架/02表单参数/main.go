package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/from",func(c *gin.Context) {
		type1 := c.DefaultPostForm("type","alert")
		username := c.PostForm("username")
		password := c.PostForm("password")
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK,fmt.Sprintf("type is %s,name is %s,password is %s,hobby is %v",type1,username,password,hobbys))
	})
	r.Run(":8080")
}
