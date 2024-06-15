package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wijohnst/templ-htmx/views"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		views.Index().Render(c, c.Writer)
	})

	router.Run(":2621")
}
