package service

import (
	"context"

	"oauth2-server/internal/model"
)

type ServerServiceInterface interface {
	GetKeys(ctx context.Context, clientId string) []string
	CreateToken(ctx context.Context, reqData model.TokenDTO) (string, error)
	ValidAuthData(data []string) bool
	ValidateJWT(token, clienId string) (model.IntrospectionResponse, error)
}
