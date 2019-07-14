package website

import (
	"html/template"

	"github.com/katoozi/gin-web-site/pkg/templatetags"

	"github.com/gin-gonic/gin"
)

// RegisterTemplateFuncs Register all custom template funcs.
// call it before load html templates
func RegisterTemplateFuncs(engine *gin.Engine) {
	engine.SetFuncMap(template.FuncMap{
		"formatAsDate": templatetags.FormatAsDate,
		"intComma":     templatetags.IntComma,
	})
}
