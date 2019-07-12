package website

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes will register all the website routes
func RegisterRoutes(engine *gin.Engine) {
	// regsiter template funcs
	engine.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	// load html files
	engine.LoadHTMLGlob("./web/templates/components/*")
	engine.LoadHTMLGlob("./web/templates/*.html")
	//engine.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	engine.GET("/", homeHandler)
}
