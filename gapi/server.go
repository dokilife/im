package gapi

import (
	db "github.com/dokilife/im/db/sqlc"
	"github.com/dokilife/im/pb"
	"github.com/dokilife/im/util"
)

// Server serves gRPC requests for our service.
type Server struct {
	pb.UnimplementedIMServer
	config util.Config
	store  db.Store
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
