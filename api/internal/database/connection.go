package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Retrieve the database connection string from the env vars
// and connect to the database
func Connect() *sqlx.DB {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		port     = os.Getenv("POSTGRES_PORT")
		dbName   = os.Getenv("POSTGRES_DATABASE")
		sslMode  = os.Getenv("POSTGRES_SSL_MODE")
	)

	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		host, port, user, password, dbName, sslMode)

	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		log.Fatal("Could not connect to database", err)
	}

	db.MustExec("SELECT * FROM pg_user")
	fmt.Println("Database successfully connected")

	return db
}