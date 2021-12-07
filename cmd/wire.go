//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package cmd

import (
	http2 "net/http"
	"simple-server-go/api/grpc"
	"simple-server-go/api/http"
	"simple-server-go/api/http/middle"
	"simple-server-go/internal/endpoint"

	"github.com/google/wire"
	grpc2 "google.golang.org/grpc"
)

type Server struct {
	Http *http2.Server
	Grpc *grpc2.Server
}

func NewServer(http *http2.Server,
	grpc *grpc2.Server) *Server {
	return &Server{
		Grpc: grpc,
		Http: http,
	}
}

// initApp init endpoints
func InitServer() (*Server, error) {
	panic(wire.Build(
		endpoint.ProviderSet,
		grpc.NewGrpc,
		middle.NewMiddleware,
		http.NewHttp,
		NewServer,
	))
}
