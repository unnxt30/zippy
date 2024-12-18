package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unnxt30/zippy/db"
)

func main(){
	r := gin.Default()
	db.RedisInit()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
