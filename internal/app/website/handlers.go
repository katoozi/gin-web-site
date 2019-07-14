package website

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/katoozi/gin-web-site/internal/pkg/models"
)

// DbCon is the sqlx db connection
var DbCon *sqlx.DB

func homeHandler(c *gin.Context) {
	usersData := []models.User{}
	err := DbCon.Select(&usersData, `SELECT "id","email","username","password","last_login","first_name","last_name" FROM "user" ORDER BY "id" ASC;`)
	if err != nil {
		log.Fatalf("Error while unmarshal to struct users data: %v", err)
	}
	// fmt.Println(sqltools.GenerateInsertQuery("user", usersData[0]))
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title":  "My First Gin Website",
		"time":   time.Now(),
		"number": 12000,
		"users":  usersData,
	})
}

func insertDataHandler(c *gin.Context) {
	usersData := []*models.User{
		models.NewUser("mohammad", "katoozi", "katoozi", "k2527806@gmail.com", "12345", time.Date(2019, 07, 11, 11, 30, 30, 0, time.UTC)),
		models.NewUser("mohammad", "katoozi", "katoozi1", "k2527806@gmail1.com", "123467", time.Date(2019, 07, 12, 12, 30, 30, 0, time.UTC)),
		models.NewUser("mohammad", "katoozi", "katoozi2", "k2527806@gmail2.com", "12346789", time.Date(2019, 07, 13, 13, 30, 30, 0, time.UTC)),
	}
	tx := DbCon.MustBegin()
	for _, user := range usersData {
		_, err := tx.Exec(user.GenerateInsertQuery())
		fmt.Println(err)
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"inserted": "true",
	})
}

// Login is post data schema
type Login struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func checkLogin(c *gin.Context) {
	// fmt.Println(c.Request.FormValue("username"))
	var loginData Login
	if err := c.ShouldBind(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.GetUser(loginData.Username, DbCon)
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
