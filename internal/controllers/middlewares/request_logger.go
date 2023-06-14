package middlewares

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			body []byte
			err  error
		)
		if c.Request.Body != nil {
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Error("could not read the request body", zap.Error(err))
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		logger.Debug("incoming request",
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.URL.Path),
			zap.String("body", string(body)),
		)

		c.Next()
	}
}
