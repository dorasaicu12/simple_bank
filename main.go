package main

import (
	"database/sql"
	"log"

	"github.com/dorasaicu12/simplebank/api"
	db "github.com/dorasaicu12/simplebank/db/sqlc"
	"github.com/dorasaicu12/simplebank/util"
	_ "github.com/lib/pq"
)

const ()

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("can connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config,store)
	if err != nil {
		log.Fatal("server err:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("server err:", err)
	}
}
