package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var downloads = Downloads{}

func main() {
	downloads.Init()
	router := gin.New()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET",
		RequestHeaders:  "Origin, Authorization, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Credentials",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.GET("/api/downloads", findInArange)
	router.Run()
}

//findInArange is the handler of /api/v1/find route.
func findInArange(c *gin.Context) {
	d := downloads.Find(c.Query("from"), c.Query("to"))
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "OK", "downloads": d})
}
