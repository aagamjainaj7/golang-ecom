package main

import (
	"database/sql"
	"log"

	"github.com/aagamjainaj7/golang-ecom/cmd/api"
	"github.com/aagamjainaj7/golang-ecom/configs"
	"github.com/aagamjainaj7/golang-ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Println("Error connecting to DB")
		log.Fatal(err)
	}
	log.Println("DB connected successfully")

}
