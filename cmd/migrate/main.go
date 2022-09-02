package main

import (
	"log"

	datastore "basic-crud/internal/datastore"
	db "basic-crud/internal/db/postgres"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		return
	}

	dbConn, _ := db.PGConnection()

	err = dbConn.AutoMigrate(&datastore.Product{})
	if err != nil {
		log.Println(err)
	}

	err = dbConn.AutoMigrate(&datastore.Role{})
	if err != nil {
		log.Println(err)
	}

	err = dbConn.AutoMigrate(&datastore.User{})
	if err != nil {
		log.Println(err)
	}
}
