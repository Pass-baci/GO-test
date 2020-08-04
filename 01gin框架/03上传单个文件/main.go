package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file,file.Filename)
		c.String(http.StatusOK,fmt.Sprintf("%s upload!",file.Filename))
	})
	r.Run(":8080")
}
