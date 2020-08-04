package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/long_async", func(c *gin.Context) {
		copyContext := c.Copy()
		go func() {
			//time.Sleep(3 * time.Second)
			log.Println("异步执行"+copyContext.Request.URL.Path)
		}()
	})
	r.GET("/long_sync", func(c *gin.Context) {
		log.Println("同步处理"+c.Request.URL.Path)
	})
	r.Run(":8080")

}
