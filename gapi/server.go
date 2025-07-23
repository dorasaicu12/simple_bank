package gapi

import (
	"fmt"

	db "github.com/dorasaicu12/simplebank/db/sqlc"
	"github.com/dorasaicu12/simplebank/pb"
	"github.com/dorasaicu12/simplebank/token"
	"github.com/dorasaicu12/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

//newServer create a new gRPC
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmectricKey)
	if err != nil {
		return nil, fmt.Errorf("can not create token maker:%v", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}


	return server, nil
}