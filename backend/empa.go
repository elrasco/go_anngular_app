package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var downloads = Downloads{}

func main() {
	downloads.Init()
	router := gin.Default()
	v1 := router.Group("/api/v1/find")

	v1.GET("/", findInArange)
	router.Run()
}

//findInArange is the handler of /api/v1/find route.
func findInArange(c *gin.Context) {
	d := downloads.Find(c.Query("from"), c.Query("to"))
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "OK", "downloads": d})
}
