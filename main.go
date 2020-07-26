package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/heroku/og/og"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", response)

	router.Run(":" + port)
}

func response(c *gin.Context) {

	urlQuery := c.Query("url")
	fmt.Println(urlQuery)

	pageInfo, e := og.GetPageInfoFromUrl(urlQuery)

	if e != nil {
		c.String(http.StatusBadRequest, "error", e)
	}

	c.JSON(http.StatusOK, pageInfo)
}
