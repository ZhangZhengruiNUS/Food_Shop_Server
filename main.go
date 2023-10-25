package main

import (
	db "Food_Shop_Server/db/sqlc"
	handler "Food_Shop_Server/handlers"
	"Food_Shop_Server/util"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config")
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Println("cannot connect to db")
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := handler.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
