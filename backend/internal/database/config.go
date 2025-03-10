package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	username string
	password string
	host     string
	port     string
	database string
}

func (c DBConfig) isEmpty() bool {
	return c.username == "" &&
		c.password == "" &&
		c.host == "" &&
		c.port == "" &&
		c.database == ""
}

func NewDBConfig() (*DBConfig, error) {
	dbConfig := &DBConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_DATABASE"),
	}
	if notEmpty := !dbConfig.isEmpty(); notEmpty {
		return dbConfig, nil
	}

	err := godotenv.Load("../.env")
	if err != nil {
		return dbConfig, fmt.Errorf("load env file failed")
	}
	return dbConfig, nil
}

func ConnectDB(config DBConfig) (*sql.DB, error) {
	connectionStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.username,
		config.password,
		config.host,
		config.port,
		config.database,
	)

	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
