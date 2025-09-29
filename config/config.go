package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DBUserName   string
	DBPassword   string
	DBHost       string
	DBPort       int
	DBName       string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failedto load the env variables: ", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required!")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required!")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is required!")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required!")
		os.Exit(1)
	}

	dbUserName := os.Getenv("DB_USER_NAME")
	if dbUserName == "" {
		fmt.Println("DB user name is required!")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB password is required!")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB host is required!")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	dbPortInt, _ := strconv.Atoi(dbPort)
	if dbPort == "" {
		fmt.Println("DB port is required!")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB name is required!")
		os.Exit(1)
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DBUserName:   dbUserName,
		DBPassword:   dbPassword,
		DBHost:       dbHost,
		DBPort:       dbPortInt,
		DBName:       dbName,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}

	return configurations
}
