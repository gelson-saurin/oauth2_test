package model

type UriParamRequest struct {
	GrantType string  `json:"grantType" binding:"required"`
	Scope     *string `json:"scope,omitempty"`
}

type CredentialsHeaderRequest struct {
	Authorization string `json:"authorization" binding:"required"`
}
