package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/dorasaicu12/simplebank/api"
	db "github.com/dorasaicu12/simplebank/db/sqlc"
	"github.com/dorasaicu12/simplebank/gapi"
	"github.com/dorasaicu12/simplebank/pb"
	"github.com/dorasaicu12/simplebank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	 runGRPCServer(config, store)
	// runGinServer(config, store)
}

func runGRPCServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("server err:", err)
	}
	grpcSever := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcSever, server)
	reflection.Register(grpcSever)
	listener, err := net.Listen("tcp", config.GrpcServerAddress)

	if err != nil {
		log.Fatal("server err when start listener:", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcSever.Serve(listener)

	if err != nil {
		log.Fatal("server err when server:", err)
	}
}




func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("server err:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("server err:", err)
	}
}
