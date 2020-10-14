package coreframeworksdatabase

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //
)

// IDataBase ...
type IDataBase interface {
	Connect() *gorm.DB
}

type dataBase struct{}

// ConnectToDatabase ...
func (*dataBase) Connect() *gorm.DB {
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

// DataBase ...
func DataBase() IDataBase {
	return &dataBase{}
}
