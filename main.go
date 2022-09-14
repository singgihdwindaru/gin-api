package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Tech Company listing API with Golang"})
}
func main() {
	router := gin.Default()
	router.GET("/", HomepageHandler)
	router.Run()
}
