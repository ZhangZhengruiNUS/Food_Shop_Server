package main

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	handler "Red_Pocket_Game_Server/handlers"
	"Red_Pocket_Game_Server/util"
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
