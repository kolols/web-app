package conf

import (
	"log"
	"os"
	"strconv"
)

const (
	hostKey       = "RGB_HOST"
	portKey       = "RGB_PORT"
	dbHostKey     = "RGB_DB_HOST"
	dbPortKey     = "RGB_DB_PORT"
	dbNameKey     = "RGB_DB_NAME"
	dbUserKey     = "RGB_DB_USER"
	dbPasswordKey = "RGB_DB_PASSWORD"
	jwtSecretKey  = "RGB_JWT_SECRET"
)

type Config struct {
	Host       string
	Port       string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	JwtSecret  string
}

func NewConfig() Config {
	host, ok := os.LookupEnv(hostKey)
	if !ok || host == "" {
		logAndPanic(hostKey)
	}

	port, ok := os.LookupEnv(portKey)
	if !ok || port == "" {
		if _, err := strconv.Atoi(port); err != nil {
			logAndPanic(portKey)
		}
	}

	dbHost, ok := os.LookupEnv(dbHostKey)
	if !ok || dbHost == "" {
		logAndPanic(dbHostKey)
	}

	dbPort, ok := os.LookupEnv(dbPortKey)
	if !ok || dbPort == "" {
		if _, err := strconv.Atoi(dbPort); err != nil {
			logAndPanic(dbPortKey)
		}
	}

	dbName, ok := os.LookupEnv(dbNameKey)
	if !ok || dbName == "" {
		logAndPanic(dbNameKey)
	}

	dbUser, ok := os.LookupEnv(dbUserKey)
	if !ok || dbUser == "" {
		logAndPanic(dbUserKey)
	}

	dbPassword, ok := os.LookupEnv(dbPasswordKey)
	if !ok || dbPassword == "" {
		logAndPanic(dbPasswordKey)
	}

	jwtSecret, ok := os.LookupEnv(jwtSecretKey)
	if !ok || jwtSecret == "" {
		logAndPanic(jwtSecretKey)
	}

	return Config{
		Host:       host,
		Port:       port,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		JwtSecret:  jwtSecret,
	}
}

func logAndPanic(envVar string) {
	log.Println("ENV variable not set or value not valid: ", envVar)
	panic(envVar)
}
