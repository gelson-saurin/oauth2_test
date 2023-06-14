package middlewares

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogResponseBody(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{body: new(bytes.Buffer), ResponseWriter: c.Writer}
		c.Writer = w
		c.Next()
		if w.body.Len() == 0 {
			logger.Info("[RESPONSE]:",
				zap.Int("status code", c.Writer.Status()),
				zap.String("body", "empty"),
			)
			return
		}
		logger.Info("[RESPONSE]:",
			zap.Int("status code", c.Writer.Status()),
			zap.String("body", w.body.String()),
		)
	}
}
