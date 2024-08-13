package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	errEnv:= godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file'")
	}

	dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

	cfg := mysql.Config{
        User:   dbUser,
        Passwd: dbPass,
        Net:    "tcp",
        Addr:   fmt.Sprintf("%s:%s", dbHost, dbPort),
        DBName: dbName,
    }

	db, err := sql.Open("mysql", cfg.FormatDSN())
	
	if err != nil {
        log.Fatal(err)
    }


	fmt.Println("Database connected")

	return db
}