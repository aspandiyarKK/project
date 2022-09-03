package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var version = "1.0.0"
var lastRequestTime time.Time

func main() {
	router := gin.Default()
	router.GET("/version", getVersion)
	router.GET("/time", getTime)
	router.GET("/lastVisit", getLastVisit)
	if err := router.Run("localhost:8080"); err != nil {
		return
	}

}

func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": version,
	})
	lastRequestTime = time.Now()
}
func getTime(c *gin.Context) {
	t := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"currentTime": t.Format(time.RFC3339),
	})
	lastRequestTime = time.Now()
}

func getLastVisit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"lastRequestTime": lastRequestTime.Format(time.RFC3339),
	})
}
