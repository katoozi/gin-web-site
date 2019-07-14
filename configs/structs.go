package configs

import (
	"strconv"

	"github.com/spf13/viper"
)

// ServerConfig is the config.yaml server section schema
type ServerConfig struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
}

// GetAddr generate the server connection address
func (s *ServerConfig) GetAddr() string {
	return s.Addr + ":" + strconv.Itoa(s.Port)
}

// DatabaseConfig is the config.yaml database section schema
type DatabaseConfig struct {
	DatabaseName string `mapstructure:"db_name"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"pass"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
}

// RedisConfig is the redis client config shema
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// GetAddr generate the server connection address
func (s *RedisConfig) GetAddr() string {
	return s.Addr + ":" + strconv.Itoa(s.Port)
}

// SetDefaultValues set the default values of server and database sections
func SetDefaultValues() {
	viper.SetDefault("database", DatabaseConfig{
		DatabaseName: "mydb",
		User:         "myuser",
		Password:     "1234",
		Host:         "localhost",
		Port:         5432,
	})
	viper.SetDefault("server", ServerConfig{
		Addr: "localhost",
		Port: 8081,
	})
}
