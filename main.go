package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var (
	Router *gin.Engine
)

func main() {
	startApp()
}

func startApp() {
	configureRouter()

	Router.GET("/ping", ping)
	Router.Run(":8080")
}

func configureRouter() {
	Router = gin.Default()
	Router.Use(gin.RecoveryWithWriter(os.Stdout))
	Router.RedirectFixedPath = false
	Router.RedirectTrailingSlash = false
	Router.HandleMethodNotAllowed = true
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
