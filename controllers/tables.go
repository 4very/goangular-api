package controllers

import (
	"log"

	"github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
	st "github.com/sommea/goangular-api/structs"
)

// Create User Table
func CreateTables(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&st.Player{}, opts)
	if createError != nil {
		log.Printf("Error while creating Player table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Player Table created")

	createError = db.CreateTable(&st.Guild{}, opts)
	if createError != nil {
		log.Printf("Error while creating Guild table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Guild Table created")

	createError = db.CreateTable(&st.Leaderboard{}, opts)
	if createError != nil {
		log.Printf("Error while creating Leaderboard table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Leaderboard Table created")

	return nil
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}
