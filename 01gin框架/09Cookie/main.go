package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NoSet"
			c.SetCookie("key_cookie", "123",60,"/cookie", "localhost", false,true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})
	r.Run(":8080")

}
