package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var downloads = Downloads{}
var corsConfig = cors.Config{
	Origins:         "*",
	Methods:         "GET",
	RequestHeaders:  "Origin, Authorization, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Credentials",
	ExposedHeaders:  "",
	MaxAge:          50 * time.Second,
	Credentials:     true,
	ValidateHeaders: false,
}

func main() {
	downloads.Init()
	router := gin.New()

	router.Use(cors.Middleware(corsConfig))

	router.GET("/api/downloads", findInArange)
	router.GET("/api/statistics", getStatistics)
	router.Run()
}

//findInArange is the handler of /api/downloads route.
func findInArange(c *gin.Context) {
	d := downloads.Find(c.Query("from"), c.Query("to"))
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "OK", "downloads": d})
}

func getStatistics(c *gin.Context) {
	s := downloads.getStatistics(c.Query("from"), c.Query("to"))
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "OK", "statistics": s})
}
