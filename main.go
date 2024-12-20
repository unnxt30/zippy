package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unnxt30/zippy/db"
	"github.com/unnxt30/zippy/frontend"
	"github.com/unnxt30/zippy/hash"
)

type URLRequest struct {
	LongURL string `json:"long_url"` // Use binding to validate
}

func main() {
	rc := db.RedisClient{}
	rc.Init()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		component := frontend.Index()
		err := component.Render(context.Background(), c.Writer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not render component"})
			return
		}
	})

	r.POST("/", func(c *gin.Context) {
		gotURL := c.Request.FormValue("long_url")
		if !hash.ValidURL(gotURL) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid URL"})
			return
		}

		shortURL, err := rc.RedisSet(gotURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not set value in redis"})
			return
		}

		hostName := c.Request.Host
		responseURL := fmt.Sprintf("http://%s/%s", hostName, shortURL)

		c.String(http.StatusOK, responseURL)
		// c.JSON(http.StatusOK, gin.H{
		// 	"long_url": gotURL,
		// 	"short_url": fmt.Sprintf("http://%s/%s",hostName, shortURL),
		// })
	})

	r.GET("/:shortURL", func(c *gin.Context) {
		shawty := c.Param("shortURL")
		longURL, err := rc.RedisGet(shawty)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "URL not found"})
			return
		}

		c.Header("Location", longURL)
		c.JSON(http.StatusFound, gin.H{
			"long_url": longURL})
	})

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
