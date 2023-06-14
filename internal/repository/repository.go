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

func NewServerRepository() ServerRepositoryInterface {
	return &ServerRepository{
		AuthorizationData: []AuthorizationData{
			{
				ClientId:     "a0897e6d0ea94f589c38278bca4e9342",
				ClientSecret: "c94dbd582d594e8aa04934f9c7ef0f52",
			},
		},
		Keys: nil,
		JWT:  nil,
	}
}

func (r *ServerRepository) GetKeys(clientId string) []rsa.PublicKey {
	var keyByUserId []rsa.PublicKey
	if r.Keys != nil {
		for _, val := range r.Keys {
			if val.ClientId == clientId {
				keyByUserId = append(keyByUserId, val.PrivateKey.PublicKey)
			}
		}
	}
	return keyByUserId
}

func (r *ServerRepository) CreateToken(clientId, token string) {
	newToken := JWT{
		ClientId: clientId,
		Token:    token,
	}
	r.JWT = append(r.JWT, newToken)
}

func (r *ServerRepository) AddKeys(clientId string, key rsa.PrivateKey) {
	newKey := Keys{
		ClientId:   clientId,
		PrivateKey: &key,
	}
	r.Keys = append(r.Keys, newKey)
}

func (r *ServerRepository) GetAuthData() []AuthorizationData {
	return r.AuthorizationData
}
