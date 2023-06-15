package repository

import "crypto/rsa"

type ServerRepository struct {
	AuthorizationData []AuthorizationData
	Keys              []Keys
	JWT               []JWT
}

type AuthorizationData struct {
	ClientId     string
	ClientSecret string
}

type Keys struct {
	ClientId   string
	PrivateKey *rsa.PrivateKey
}

type JWT struct {
	ClientId string
	Token    string
}
