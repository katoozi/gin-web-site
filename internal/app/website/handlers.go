package website

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/katoozi/gin-web-site/internal/pkg/auth"
)

// DbCon is the sqlx db connection
var DbCon *sqlx.DB

// RedisCon connection
var RedisCon *redis.Client

func homeHandler(c *gin.Context) {
	usersData := []auth.User{}
	err := DbCon.Select(&usersData, `SELECT "id","email","username","password","last_login","first_name","last_name","is_active" FROM "user" ORDER BY "id" ASC;`)
	if err != nil {
		log.Fatalf("Error while unmarshal to struct users data: %v", err)
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title":  "My First Gin Website",
		"time":   time.Now(),
		"number": 12000,
		"users":  usersData,
	})
}

func insertDataHandler(c *gin.Context) {
	usersData := []*auth.User{
		auth.NewUser("mohammad", "katoozi", "katoozi", "k2527806@gmail.com", "12345"),
		auth.NewUser("mohammad", "katoozi", "katoozi1", "k2527806@gmail1.com", "123467"),
		auth.NewUser("mohammad", "katoozi", "katoozi2", "k2527806@gmail2.com", "12346789"),
	}
	tx := DbCon.MustBegin()
	for _, user := range usersData {
		tx.Exec(user.GenerateInsertQuery())
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"inserted": "true",
	})
}

func checkLogin(c *gin.Context) {
	var loginData Login
	if err := c.ShouldBind(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := auth.GetUser(loginData.Username, DbCon)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	err := user.Compare(loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	c.SetCookie("user", user.Password, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": user.Password,
	})
}

func testFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "you are logged in.",
	})
}

func createCookie(c *gin.Context) {
	c.SetCookie("sessionid", "s4qtx20gmkixl3yz6fqzbdecvlyj3smt", 3600, "/", "127.0.0.1", false, false)
	c.JSON(http.StatusOK, gin.H{
		"status": "cookie set successfully",
	})
}
