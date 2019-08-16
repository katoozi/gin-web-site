package website

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes will register all the website routes
func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/", homeHandler)
	engine.GET("/insert-data", insertDataHandler)
	engine.POST("/login", checkLogin)
	engine.GET("/create-cookie", createCookie)
	engine.GET("/ws", func(c *gin.Context) {
		serveWs(c.Writer, c.Request)
	})

	authorized := engine.Group("/user")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/read", testFunc)
	}
}
