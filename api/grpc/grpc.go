package grpc

import (
	"google.golang.org/grpc"
)

func NewGrpc() (s *grpc.Server, err error) {
	s = grpc.NewServer()
	return
}
