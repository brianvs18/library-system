package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *sql.DB

func Connection() error {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbName := os.Getenv("db_name")
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_password")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("connection error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("ping error: %v", err)
	}

	DB = db
	return nil
}
