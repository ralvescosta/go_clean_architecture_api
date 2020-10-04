package core

import (
	migrations "gomux_gorm/src/core/database/table_models"
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

	connection.AutoMigrate(&migrations.Users{}, &migrations.Permissions{}, &migrations.UsersPermissions{}, &migrations.Sessions{})

	database := connection.DB()
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return connection
}
