
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
  _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	DbHostSurvey string
	DbSurvey     *gorm.DB
	Port         string
	LogMode      bool
	Env          string
)

type DatabaseConfig struct {
	Db       string `json:"db"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Environment  string `json:"environment"`
	Port         string `json:"port"`
	LogMode      bool   `json:"logMode"`
	MySql         []DatabaseConfig
}

func InitConfig() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	databaseConfig := []DatabaseConfig{
		{
			Db:       os.Getenv("MYSQL_DB"),
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
		},
	}

	logMode, err := strconv.ParseBool(os.Getenv("LOG_MODE"))
	if err != nil {
		fmt.Println("Failed to parse LOG_MODE")
		os.Exit(2)
	}

	c := Config{
		Environment:  os.Getenv("ENVIRONMENT"),
		Port:         os.Getenv("PORT"),
		LogMode:      logMode,
		MySql:         databaseConfig,
	}

	// Set service domain and application port
	Env = c.Environment
	Port = c.Port
	LogMode = c.LogMode
	InitDb(c)
}
