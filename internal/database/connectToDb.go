package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDb() {
	env_err := godotenv.Load()
	if env_err != nil {
		panic("error env")
	}

	var err error

	connStr := "postgres://" + os.Getenv("dbuser") + ":" + os.Getenv("dbpassword") + "@" + os.Getenv("dbhost") + ":5432/" + os.Getenv("dbname") + "?sslmode=disable"
	DB, err = sql.Open("pgx", connStr)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}
}
