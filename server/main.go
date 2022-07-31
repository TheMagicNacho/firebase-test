package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "server running")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
