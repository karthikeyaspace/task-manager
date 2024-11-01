package config

import "fmt"

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
	Port   string
}

func NewConfig() *Config {
	return &Config{
		DBUser: "api",
		DBPass: "api",
		DBHost: "localhost",
		DBPort: "5432",
		DBName: "api",
		Port:   ":8080",
	}
}

func (c *Config) GetDBURI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)

}
