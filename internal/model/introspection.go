package model

type TokenIntrospectionRequest struct {
	Token string `json:"token" binding:"required"`
}

type IntrospectionResponse struct {
	Active   bool    `json:"active" binding:"required"`
	ClientId *string `json:"clientId,omitempty"`
}
