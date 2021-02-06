package v1

import (
	"context"
	"database/sql"

	v1 "github.com/goblin2018/microservice_demo/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type todoServiceServer struct {
	db *sql.DB
}

func NewTodoServiceServer(db *sql.DB) v1.ToDoServiceServer {
	return &todoServiceServer{db: db}
}

func (s *todoServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented, "unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *todoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "faied to connect to database-> "+err.Error())
	}
	return c, nil
}
