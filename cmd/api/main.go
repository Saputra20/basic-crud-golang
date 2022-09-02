package main

import (
	db "basic-crud/internal/db/postgres"
	"basic-crud/internal/server"
	"log"
	"os"
)

func main() {
	if os.Getenv("GO_ENV") == "development" {
		log.SetFlags(log.LstdFlags | log.Llongfile)
	}

	dbConn, err := db.PGConnection()

	if err != nil || dbConn == nil {
		log.Println("Error with postgre DB", err)
		return
	}

	s := server.NewServer()

	s.Run()
}
