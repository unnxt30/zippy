package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unnxt30/zippy/db"
)

type URLRequest struct {
	LongURL string `json:"long_url"` // Use binding to validate
}

func main(){
	rc := db.RedisClient{}
	rc.Init()
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		var req URLRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}		

		shortURL, err := rc.RedisSet(req.LongURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not set value in redis"})
				return 
			}


		c.JSON(http.StatusOK, gin.H{
			"long_url": req.LongURL,
			"short_url": fmt.Sprintf("http://localhost/%s", shortURL),
		})
	})
	r.Run()
}
