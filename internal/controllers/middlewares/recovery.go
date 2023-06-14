package middlewares

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			logger.Info("recovery from panic:", zap.Error(err.(error)))
			debug.PrintStack()
		}()

		c.Next()
	}
}
