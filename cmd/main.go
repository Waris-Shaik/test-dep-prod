package main

import (
	"database/sql"
	"log"
	"test-dep-prod/cmd/api"
	"test-dep-prod/configs"
	"test-dep-prod/db"

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
		log.Fatalf("could not establish connecttion: %v", err)
	}

	initStorage(db)
	// server-instance
	server := api.NewAPIServer(":8000", db)
	if err := server.Run(); err != nil {
		log.Fatal("could not start server", err)
	}
}
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("error while pinging database:", err)
	}
	log.Printf("database connected to host:%v successfully🔥🔥🔥", configs.Envs.DBAddress)
}
