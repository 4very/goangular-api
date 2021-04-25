package config

import (
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
	controllers "github.com/sommea/goangular-api/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	godotenv.Load()

	opts := &pg.Options{
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     "192.168.0.201:5432",
		Database: os.Getenv("DB_DATABASE"),
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")

	controllers.CreateTables(db)
	controllers.InitiateDB(db)
	return db
}
