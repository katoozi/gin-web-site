package website

import (
	"fmt"
	"html/template"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

// RegisterTemplateFuncs Register all custom template funcs.
// call it before load html templates
func RegisterTemplateFuncs(engine *gin.Engine) {
	engine.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"intComma":     commaSeperated,
	})
}

func commaSeperated(number int) string {
	return fmt.Sprintf("%s", humanize.Comma(int64(number)))
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", year, month, day)
}
