package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/search", searchHandler)
	router.Run("localhost:8000")
}
