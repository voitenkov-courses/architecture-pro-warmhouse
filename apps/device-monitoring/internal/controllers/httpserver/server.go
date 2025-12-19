package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "github.com/voitenkov-courses/architecture-pro-warmhouse/apps/device-monitoring/api/telemetry-access"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ api.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// Удаление телеметрии устройства
// (DELETE /telemetry/{serialNo})
func (Server) DeleteTelemetry(c *gin.Context, serialNo string, params api.DeleteTelemetryParams) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Получение телеметрии устройства
// (GET /telemetry/{serialNo})
func (Server) GetTelemetry(c *gin.Context, serialNo string, params api.GetTelemetryParams) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}
