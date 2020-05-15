package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func getDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		"host="+os.Getenv("HOSTNAME")+
			" port="+os.Getenv("DB_PORT")+
			" user="+os.Getenv("USER")+
			" sslmode=disable"+
			" dbname="+os.Getenv("DBNAME")+
			" password="+os.Getenv("PASSWORD"),
	)
	if err != nil {
		log.Printf("Error at open Database: %v", err)
		return nil, err
	}

	return db, nil
}
