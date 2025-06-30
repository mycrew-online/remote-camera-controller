package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// GinLoggerMiddleware returns a Gin middleware that logs requests using the SimConnectManager's logger.
// GinLoggerMiddleware returns a Gin middleware that logs requests using the SimConnectManager's logger.
func GinLoggerMiddleware(mgr *manager.SimConnectManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		status := c.Writer.Status()
		msg := fmt.Sprintf("[GinRouteHandler] %s %s - %d", method, path, status)
		mgr.Logger().Info(msg)
	}
}
