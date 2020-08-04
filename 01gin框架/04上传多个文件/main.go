package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("get err %s",err))
		}
		files := form.File["file"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file,file.Filename); err != nil {
				c.String(http.StatusOK,fmt.Sprintf("upload err %s",err.Error()))
				return
			}
			log.Println(file.Filename)
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})
	r.Run(":8080")
}
