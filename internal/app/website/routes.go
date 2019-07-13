package website

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes will register all the website routes
func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/", homeHandler)
}
