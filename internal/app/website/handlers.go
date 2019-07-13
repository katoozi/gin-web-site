package website

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	usersData := []User{}
	err := dbCon.Select(&usersData, `SELECT "id","email","username","password","last_login","first_name","last_name" FROM "user" ORDER BY "id" ASC;`)
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
	usersData := []*User{
		NewUser("mohammad", "katoozi", "katoozi", "k2527806@gmail.com", "12345", "2019-07-13"),
		NewUser("mohammad", "katoozi", "katoozi1", "k2527806@gmail1.com", "123467", "2019-07-14"),
		NewUser("mohammad", "katoozi", "katoozi2", "k2527806@gmail2.com", "12346789", "2019-07-15"),
	}
	tx := dbCon.MustBegin()
	for _, user := range usersData {
		tx.Exec(user.GenerateInsertQuery())
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
	user := GetUser(loginData.Username)
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
