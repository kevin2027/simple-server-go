package handle

import (
	"simple-server-go/api/http/middle"
	"simple-server-go/internal/endpoint"

	"github.com/gin-gonic/gin"
)

func HealthBluePrint(v1 *gin.RouterGroup, e *endpoint.Endpoints, m *middle.Middleware) {
	r := v1.Group("/health")
	// 小程序接口
	r.GET("", checkHealth(e))
}

func checkHealth(e *endpoint.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatus(200)
	}

}
