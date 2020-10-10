package coreframeworksdatabase

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //
)

// ConnectToDatabase ...
func ConnectToDatabase() *gorm.DB {
	connection, err := gorm.Open("postgres", "user=postgres password=12345 dbname=default sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	database := connection.DB()
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return connection
}
