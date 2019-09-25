package website

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/katoozi/gin-web-site/pkg/templatefuncs"
	"github.com/streadway/amqp"
)

// Initial will set the redis and postgres connections, register all routes, load templates,
// enable middleswares, upgrade connection to websocket and ...
func Initial(redis *redis.Client, postgres *sqlx.DB, rabbitmq *amqp.Channel) *gin.Engine {
	// set the connections
	RedisCon = redis
	DbCon = postgres
	RabbitMQCon = rabbitmq

	// config the gin framework
	r := gin.Default()

	// serve static files
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
	// enable AuthMiddleware for /read endpoint
	authorized.Use(AuthRequired())
	{
		authorized.GET("/read", testFunc)
	}

	return r
}
