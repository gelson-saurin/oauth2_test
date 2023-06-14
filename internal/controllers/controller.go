package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"oauth2-server/internal/controllers/path"
	"oauth2-server/internal/service"
)

type ServerConfigController struct {
	service service.ServerServiceInterface
	engine  *gin.Engine
	logger  *zap.Logger
}

func NewServerConfigController(e *gin.Engine, s service.ServerServiceInterface, l *zap.Logger) {
	endpointConfigCtrl := &ServerConfigController{
		service: s,
		engine:  e,
		logger:  l,
	}
	endpointConfigCtrl.setUpRoutes()
}

func (c *ServerConfigController) GetKeys(ctx *gin.Context) {
	rawDecodedText, err := c.validateAuth(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	decodedText := strings.Split(string(rawDecodedText), ":")

	if !c.checkAuthData(decodedText) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	data := c.service.GetKeys(ctx.Request.Context(), decodedText[0])
	if len(data) == 0 {
		c.logger.Debug(fmt.Sprintf("no found keys to cliendId: %v", decodedText[0]))
		ctx.AbortWithStatus(http.StatusOK)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (c *ServerConfigController) CreateToken(ctx *gin.Context) {
	var uriParams UriParamRequest
	errUri := ctx.ShouldBindUri(uriParams)
	if errUri != nil {
		c.logger.Debug(fmt.Sprintf("failed to bind URI, error: %v", errUri))
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	rawDecodedText, err := c.validateAuth(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	decodedText := strings.Split(string(rawDecodedText), ":")

	if !c.checkAuthData(decodedText) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	reqData := TokenDTO{
		ClientId:     decodedText[0],
		ClientSecret: decodedText[1],
		GrantType:    uriParams.GrantType,
	}

	if uriParams.Scope != nil {
		reqData.Scope = uriParams.Scope
	}

	data, err := c.service.CreateToken(ctx.Request.Context(), reqData)
	if err != nil {
		c.logger.Error("INSERT", zap.Error(err))
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	c.logger.Info(fmt.Sprintf("successful token creation, response: %v", data))
	ctx.JSON(http.StatusOK, gin.H{"token": data})
}

func (c *ServerConfigController) validateAuth(ctx *gin.Context) ([]byte, error) {
	var requestCredentials CredentialsHeaderRequest
	errHeader := ctx.ShouldBindHeader(requestCredentials)
	if errHeader != nil {
		c.logger.Debug(fmt.Sprintf("failed to bind header's autorization, error: %v", errHeader))
		return nil, fmt.Errorf("failed to bind header's autorization, status: %v", http.StatusBadRequest)
	}

	rawDecodedText, err := base64.StdEncoding.DecodeString(requestCredentials.Authorization)
	if err != nil {
		c.logger.Debug(fmt.Sprintf("failed to decode autorization, error: %v", err))
		return nil, fmt.Errorf("failed to decode autorization, status: %v", http.StatusBadRequest)
	}

	return rawDecodedText, nil
}

func (c *ServerConfigController) checkAuthData(authData []string) bool {
	return c.service.ValidAuthData(authData)
}

func (c ServerConfigController) setUpRoutes() {
	c.engine.GET(path.GetKeys, c.GetKeys)
	c.engine.POST(path.CreateToken, c.CreateToken)
}
