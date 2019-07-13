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
	tx := dbCon.MustBegin()
	tx.MustExec(`INSERT INTO "user" (first_name, last_name, email, username, password, last_login) VALUES ($1, $2, $3, $4, $5, $6)`, "mohammad", "katoozi", "k2527806@gmail.com", "katoozi", "1234", "2019-07-13")
	tx.MustExec(`INSERT INTO "user" (first_name, last_name, email, username, password, last_login) VALUES ($1, $2, $3, $4, $5, $6)`, "mohammad", "katoozi", "k2527806@gmail1.com", "katoozi1", "1234", "2019-07-13")
	tx.MustExec(`INSERT INTO "user" (first_name, last_name, email, username, password, last_login) VALUES ($1, $2, $3, $4, $5, $6)`, "mohammad", "katoozi", "k2527806@gmail2.com", "katoozi2", "1234", "2019-07-13")
	// tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	// tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	// tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	// tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"inserted": "true",
	})
}
