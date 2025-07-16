package api

import (
	"fmt"

	db "github.com/dorasaicu12/simplebank/db/sqlc"
	"github.com/dorasaicu12/simplebank/token"
	"github.com/dorasaicu12/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	//add route
	server.setupRoute()

	return server, nil
}

func (s *Server) setupRoute() *gin.Engine {
	router := gin.Default()

	router.POST("/users", s.createUser)
	router.POST("/users/login", s.LoginUser)

	router.GET("/yeu-em", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte("<h1 style=\"text-align:center;\">Yêu em nhiều</h1>"))
	})


	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))
	authRoutes.POST("/account", s.createAccount)
	authRoutes.GET("/account/:id", s.getAccount)
	authRoutes.GET("/accounts", s.listAccount)

	authRoutes.POST("/tranfers", s.CreateTranfer)

	s.router = router
	return router
}

func (s *Server) Start(addres string) error {
	return s.router.Run(addres)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
