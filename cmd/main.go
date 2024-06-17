package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wijohnst/templ-htmx/internal"
)

func main() {
	router := gin.Default()

	app := internal.Providers{Router: router}

	app.Routes()

	router.Run(":2621")
}
