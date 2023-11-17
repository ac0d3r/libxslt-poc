package main

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed poc.svg
var Poc string

func main() {
	r := gin.New()
	r.StaticFile("/", "./index.svg")

	r.GET("/poc", func(c *gin.Context) {
		c.Header("Cross-Origin-Resource-Sharing", "*")
		c.Data(http.StatusOK, "image/svg+xml", []byte(Poc))
	})
	r.Run(":8080")
}
