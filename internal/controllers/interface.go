package controller

import "github.com/gin-gonic/gin"

type ServerControllerInterface interface {
	GetKeys(ctx *gin.Context)
	CreateToken(ctx *gin.Context)
}
