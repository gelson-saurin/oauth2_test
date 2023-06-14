package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"

	controller "oauth2-server/internal/controllers"
	"oauth2-server/internal/repository"

	"github.com/golang-jwt/jwt"
)

var Keys *rsa.PrivateKey

type ServerService struct {
	serverRepository repository.ServerRepositoryInterface
	logger           *zap.Logger
}

func NewEndpointConfigService(r repository.ServerRepositoryInterface, l *zap.Logger) ServerServiceInterface {
	return &ServerService{
		serverRepository: r,
		logger:           l,
	}
}

func (ecs *ServerService) GetKeys(ctx context.Context, clientId string) []string {
	ecs.logger.Debug("start GET public keys by clientId")
	var publicKeySet []string

	data := ecs.serverRepository.GetKeys(clientId)

	for _, value := range data {
		keyItem, err := json.MarshalIndent(value, "", "")
		if err != nil {
			ecs.logger.Error(fmt.Sprintf("unable to marshal public keys: %s", err))
		}
		publicKeySet = append(publicKeySet, string(keyItem[:]))
	}
	return publicKeySet
}

func (ecs *ServerService) CreateToken(ctx context.Context, token controller.TokenDTO) (string, error) {
	ecs.logger.Debug("start Create JWT token")

	tokenValue, err := generateJWT(token.ClientId)
	if err != nil {
		return "", err
	}

	ecs.serverRepository.AddKeys(token.ClientId, *Keys)
	ecs.serverRepository.CreateToken(token.ClientId, tokenValue)
	return tokenValue, nil
}

func (ecs *ServerService) ValidateJWT(t, clientId string) (controller.IntrospectionResponse, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unauthorized")
		}
		//
		//kid, _ := token.Header["kid"]
		return Keys.Public(), nil
	})

	claims := token.Claims.(jwt.MapClaims)
	if claims["user"] != clientId {
		return controller.IntrospectionResponse{}, fmt.Errorf("unauthorized")
	}

	if err != nil {
		return controller.IntrospectionResponse{}, err
	}
	if !token.Valid {
		return controller.IntrospectionResponse{}, fmt.Errorf("unauthorized")
	}

	response := controller.IntrospectionResponse{
		Active:   token.Valid,
		ClientId: &clientId,
	}
	return response, nil
}

func (ecs *ServerService) ValidAuthData(data []string) bool {
	for _, val := range ecs.serverRepository.GetAuthData() {
		if val.ClientId == data[0] && val.ClientSecret == data[1] {
			return true
		}
	}
	return false
}

func generateJWT(clientId string) (string, error) {
	newToken := jwt.New(jwt.SigningMethodRS256)
	claims := newToken.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = clientId
	newToken.Header["kid"] = "RSA-256"

	var err error
	Keys, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	signToken, errSign := newToken.SignedString(Keys)
	if errSign != nil {
		return "", err
	}
	return signToken, nil
}
