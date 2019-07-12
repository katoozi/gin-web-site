package website

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
