package website

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title":  "My First Gin Website",
		"time":   time.Now(),
		"number": 12000,
	})
}
