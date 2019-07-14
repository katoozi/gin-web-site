package website

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/katoozi/gin-web-site/pkg/templatefuncs"
)

// RegisterTemplateFuncs Register all custom template funcs.
// call it before load html templates
func RegisterTemplateFuncs(engine *gin.Engine) {
	engine.SetFuncMap(template.FuncMap{
		"formatAsDate": templatefuncs.FormatAsDate,
		"intComma":     templatefuncs.IntComma,
	})
}
