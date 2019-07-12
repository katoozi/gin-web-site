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

// GetAddr generate the addr for our router
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
