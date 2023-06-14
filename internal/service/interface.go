package service

import (
	"context"

	controller "oauth2-server/internal/controllers"
)

type ServerServiceInterface interface {
	GetKeys(ctx context.Context, clientId string) []string
	CreateToken(ctx context.Context, reqData controller.TokenDTO) (string, error)
	ValidAuthData(data []string) bool
}
