package website

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthRequired will check for cookie and session data in redis.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("sessionid")
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			c.Abort()
			return
		}
		data, err := RedisCon.Get("session:" + cookie).Result()
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			c.Abort()
			return
		}
		c.Set("session_data", data)
		c.Next()
	}
}
