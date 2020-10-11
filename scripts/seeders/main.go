package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //
)

func main() {
	fmt.Println("Db connection...")
	connection, err := gorm.Open("postgres", "user=postgres password=12345 dbname=default sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	database := connection.DB()
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Run Seed...")
	connection.Exec(`
		INSERT INTO permissions
			(role, description) VALUES
			('unauthorized', 'unauthorized'),
			('user', 'user'),
			('admin', 'admin')
	`)
	fmt.Println("All Done!")

}
