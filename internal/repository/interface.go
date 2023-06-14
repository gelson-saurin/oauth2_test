package repository

import "crypto/rsa"

type ServerRepositoryInterface interface {
	GetKeys(clientId string) []rsa.PublicKey
	CreateToken(clientId, token string)
	AddKeys(clientId string, key rsa.PrivateKey)
	GetAuthData() []AuthorizationData
}
