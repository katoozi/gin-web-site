package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/katoozi/gin-web-site/internal/pkg/auth"
	"github.com/spf13/viper"

	_ "github.com/lib/pq" // register postgresql driver
)

// Init will intiate the db and load settings from yaml file
func initialPostgres() *sqlx.DB {
	// generate the postgres connect address
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.pass")
	dbName := viper.GetString("database.db.name")
	dataSourceName := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=disable",
		host+":"+port,
		user,
		dbName,
		pass,
	)

	// connect to postgresql
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Connect to db Failed: %v", err)
	}
	auth.MigrateTables(db)
	return db
}

func initialRedis() *redis.Client {
	// use redis sentinel for high availability and failover
	redisAddrs := viper.GetString("redis.sentinels")
	redisDB := viper.GetInt("redis.db")
	redisPass := viper.GetString("redis.pass")

	sentinelAddrs := strings.Split(redisAddrs, ",")
	redisClient := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: sentinelAddrs,
		Password:      redisPass,
		DB:            redisDB,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Error while connect to redis: %v\n", err)
	}
	return redisClient
}
