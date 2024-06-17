package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/wijohnst/templ-htmx/views"
)

/*
indexPageHandler()

Pointer Receiver function that adds a `gin.HandlerFun` to our Providers struct
*/
func (app *Providers) indexPageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Render(ctx, views.Index())
	}
}

func (app *Providers) noRouteProvider() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Render(ctx, views.NoRoutePage())
	}
}
