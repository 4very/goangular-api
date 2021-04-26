package controllers

import (
	"log"
	"reflect"

	"github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
	st "github.com/sommea/goangular-api/structs"
)

func getName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func createTable(db *pg.DB, m interface{}) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(m, opts)
	if createError != nil {
		log.Printf("Error while creating table "+getName(m)+", Reason: %v\n", createError)
		return createError
	}
	log.Printf(getName(m) + " Table created")
	return nil
}

// Create User Table
func CreateTables(db *pg.DB) error {
	var tables []interface{}
	tables = append(tables, &st.Guild{}, &st.Player{}, &st.ComPlayerData{}, &st.Fight{}, &st.Report{}, &st.ComData{})

	for _, elt := range tables {
		createTable(db, elt)
	}

	return nil
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}
