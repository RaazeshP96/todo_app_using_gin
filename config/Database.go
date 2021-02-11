package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	port, _ := strconv.Atoi(os.Getenv("Port"))
	dbConfig := DBConfig{
		Host:     os.Getenv("Host"),
		Port:     port,
		User:     os.Getenv("User"),
		DBName:   os.Getenv("DBName"),
		Password: os.Getenv("Password"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
