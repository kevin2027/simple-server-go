package http

import (
	"fmt"
	"log"
	"net/http"
	"simple-server-go/api/http/handle"
	"simple-server-go/api/http/middle"
	"simple-server-go/init/global"
	"simple-server-go/internal/endpoint"

	"git.yupaopao.com/ops-public/kit/middleware"
	"github.com/gin-gonic/gin"
)

func NewHttp(e *endpoint.Endpoints, m *middle.Middleware) *http.Server {
	ginEngine := gin.New()
	ginEngine.Use(middleware.Error, middleware.LoggerFormate())
	api := ginEngine.Group("/api/simple")

	handle.HealthBluePrint(api, e, m)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.Config.Server.Http.Port),
		Handler: ginEngine,
	}
	log.Printf(" [*] port:%d Waiting for connection. To exit press CTRL+C", global.Config.Server.Http.Port)
	return s
}
