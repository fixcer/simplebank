package main

import (
	"database/sql"
	"fmt"
	"github.com/fixcer/simplebank/api"
	db "github.com/fixcer/simplebank/db/sqlc"
	"github.com/fixcer/simplebank/utils"
	"log"

	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(fmt.Sprintf("%s:%d", config.ServerAddress, config.ServerPort))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
