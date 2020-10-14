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

	fmt.Println("Run Migrations...")
	connection.Exec(`
		CREATE TABLE users 
			( 
				id SERIAL NOT NULL PRIMARY KEY, 
				name varchar(255) NOT NULL,
				last_name varchar(255) NOT NULL,
				email varchar(255) NOT NULL,
				password varchar(255) NOT NULL,
				created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP WITH TIME ZONE
			);
		`)

	connection.Exec(`
		CREATE TABLE permissions 
			( 
				id SERIAL NOT NULL PRIMARY KEY, 
				role varchar(255) NOT NULL, 
				description varchar(255), 
				created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP WITH TIME ZONE
			);
		`)

	connection.Exec(`
		CREATE TABLE users_permissions
			(
				id SERIAL NOT NULL PRIMARY KEY,
				user_id int NOT NULL,
				user_name varchar(255) NOT NULL,
				user_email varchar(255) NOT NULL,
				permission_id int NOT NULL,
				permission_role varchar(255) NOT NULL,
				created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP WITH TIME ZONE,
				FOREIGN KEY (user_id) REFERENCES users(id),
				FOREIGN KEY (permission_id) REFERENCES permissions(id)
			);
		`)

	connection.Exec(`
			CREATE TABLE sessions 
				(
					id SERIAL NOT NULL PRIMARY KEY,
					user_id int NOT NULL,
					user_name varchar(255) NOT NULL,
					user_email varchar(255) NOT NULL,
					permission_id int NOT NULL,
					permission_role varchar(255) NOT NULL,
					agent varchar(255) NOT NULL,
					remote_address varchar(255) NOT NULL,
					local_address varchar(255) NOT NULL,
					local_port varchar(255) NOT NULL,
					access_token varchar(255) NOT NULL,
					created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
					deleted_at TIMESTAMP WITH TIME ZONE
				);
		`)
	fmt.Println("All Done...")
}
