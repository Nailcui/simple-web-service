package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	hostname, err := os.Hostname()
	resp := map[string]string{
		"hostname": hostname,
	}
	if err != nil {
		log.Panicln(err)
	}
	r.GET("/hostname", func(c *gin.Context) {
		c.JSON(http.StatusOK, resp)
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, resp)
	})
	err = r.Run(":8080")
	if err != nil {
		log.Panicln(err)
	}
}
