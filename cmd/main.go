package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"oauth2-server/internal/controllers"
	"oauth2-server/internal/repository"
	"oauth2-server/internal/service"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("unable to load zap logger, error: %v", err))
	}

	r := gin.Default()
	rep := repository.NewServerRepository()
	srv := service.NewEndpointConfigService(rep, logger)
	controller.NewServerConfigController(r, srv, logger)

	err = r.Run()
	if err != nil {
		panic(fmt.Errorf("unable to load gin engine, error: %v", err))
	}
}
