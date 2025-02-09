package gapi

import (
	db "github.com/jpmoraess/appointment-api/db/sqlc"
	"github.com/jpmoraess/appointment-api/internal/usecases"
	"github.com/jpmoraess/appointment-api/pb"
	"github.com/jpmoraess/appointment-api/pkg/utils"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	config   utils.Config
	store    db.Store
	register *usecases.Register
}

func NewServer(config utils.Config, store db.Store, register *usecases.Register) *Server {
	return &Server{
		config:   config,
		store:    store,
		register: register,
	}
}
