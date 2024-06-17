package internal

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// Accepts a context and writes the rendered template to the response body
func Render(ctx *gin.Context, template templ.Component) error {

	return template.Render(ctx.Request.Context(), ctx.Writer)
}
