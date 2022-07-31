package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "server running change ")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
