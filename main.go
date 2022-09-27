package main

import (
	"github.com/gin-gonic/gin"
	"github.com/banzaicloud/go-gin-prometheus"
)

func main() {
	r := gin.New()

	p := ginprometheus.NewPrometheus("gin", []string{})
	p.SetListenAddress(":8081")
	p.Use(r, "/metrics")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})

	r.Run(":8080")
}

// this gin is referred from https://github.com/banzaicloud/go-gin-prometheus
