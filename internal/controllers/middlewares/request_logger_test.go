package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func loggerCapture() (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(zapcore.DebugLevel)
	logger := zap.New(core)
	return logger, logs
}

// TestRequestLogger checks:
// - logs correctly the request information
// - the request reaches the handler
// - request body is not lost in the middleware
func TestRequestLogger(t *testing.T) {
	rec := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rec)
	logger, logs := loggerCapture()
	engine.POST("/testing", RequestLogger(logger), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	logAssertions := func(logs *observer.ObservedLogs, expectedBody string) {
		assert.NotEqual(t, 0, logs.Len())
		requestLogged := logs.FilterMessage(`incoming request`)
		entry := requestLogged.All()[requestLogged.Len()-1]
		fields := entry.ContextMap()
		assert.Equal(t, `incoming request`, entry.Message)
		method, okMethod := fields["method"]
		assert.True(t, okMethod)
		assert.Equal(t, http.MethodPost, method)
		body, okBody := fields["body"]
		assert.True(t, okBody)
		assert.Equal(t, expectedBody, body)
		uri, okUri := fields["uri"]
		assert.True(t, okUri)
		assert.Equal(t, `/testing`, uri)
	}

	t.Run("request without body", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, `/testing`, nil)
		assert.NoError(t, err)
		ctx.Request = request
		engine.HandleContext(ctx)

		logAssertions(logs, "")

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		assert.Nil(t, ctx.Request.Body)
	})

	t.Run("request with body", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, `/testing`, bytes.NewBuffer([]byte(`{i am a body}`)))
		assert.NoError(t, err)
		ctx.Request = request
		engine.HandleContext(ctx)

		logAssertions(logs, `{i am a body}`)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		assert.NotNil(t, ctx.Request.Body)
		body, errBody := io.ReadAll(ctx.Request.Body)
		assert.NoError(t, errBody)
		assert.Equal(t, `{i am a body}`, string(body))
	})
}
