package model

type UriParamRequest struct {
	GrantType string  `json:"grantType" binding:"required"`
	Scope     *string `json:"scope,omitempty"`
}

type CredentialsHeaderRequest struct {
	Authorization string `json:"authorization" binding:"required"`
}

type TokenDTO struct {
	ClientId     string  `json:"clientId" binding:"required"`
	ClientSecret string  `json:"clientSecret" binding:"required"`
	GrantType    string  `json:"grantType" binding:"required"`
	Scope        *string `json:"scope,omitempty"`
}

type TokenIntrospectionRequest struct {
	Token string `json:"token" binding:"required"`
}

type IntrospectionResponse struct {
	Active   bool    `json:"active" binding:"required"`
	ClientId *string `json:"clientId,omitempty"`
}
