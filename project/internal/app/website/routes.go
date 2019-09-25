package website

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/katoozi/gin-web-site/pkg/templatefuncs"
)

// Initial will config the package
func Initial(redis *redis.Client, postgres *sqlx.DB) *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./web/build/static")
	r.StaticFile("/manifest.json", "./web/build/manifest.json")

	// load template funcs
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": templatefuncs.FormatAsDate,
		"intComma":     templatefuncs.IntComma,
	})

	// load html files
	r.LoadHTMLGlob("./web/build/*.html")

	r.GET("/", homeHandler)
	r.GET("/insert-data", insertDataHandler)
	r.POST("/login", checkLogin)
	r.GET("/create-cookie", createCookie)
	r.GET("/ws", func(c *gin.Context) {
		// initial the websocket endpoint
		serveWs(c.Writer, c.Request)
	})

	authorized := r.Group("/user")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/read", testFunc)
	}
	return r
}
