package model

type TokenDTO struct {
	ClientId     string  `json:"clientId" binding:"required"`
	ClientSecret string  `json:"clientSecret" binding:"required"`
	GrantType    string  `json:"grantType" binding:"required"`
	Scope        *string `json:"scope,omitempty"`
}
