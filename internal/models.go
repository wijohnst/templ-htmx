package internal

import (
	"github.com/gin-gonic/gin"
)

/*
This struct defines the core functionality of the server
and will be initialized as `app` on start (see main.go)

Usage: app.router().GET
*/
type Providers struct {
	Router *gin.Engine
	// Logger *Logger // <- points to some Logger struct
}
