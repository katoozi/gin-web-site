package commands

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/katoozi/gin-web-site/internal/app/website"
	"github.com/katoozi/gin-web-site/internal/pkg/auth"
	"github.com/spf13/viper"

	_ "github.com/lib/pq" // register postgresql driver
)

var dbCon *sqlx.DB

// Init will intiate the db and load settings from yaml file
func Init() {
	// connect to postgresql
	db, err := sqlx.Connect("postgres", fetchDatabaseConfig())
	if err != nil {
		log.Fatalf("Connect to db Failed: %v", err)
	}
	dbCon = db
	auth.MigrateTables(db)
	website.DbCon = db
}

func fetchDatabaseConfig() string {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.pass")
	dbName := viper.GetString("database.db.name")
	return fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=disable",
		host+":"+port,
		user,
		dbName,
		pass,
	)

}
