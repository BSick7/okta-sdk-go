package app

type Credentials struct {
	Scheme           *string                      `json:"scheme"`
	UserNameTemplate *CredentialsUserNameTemplate `json:"userNameTemplate"`
	Signing          CredentialsSigning           `json:"signing"`
	UserName         string                       `json:"userName"`
	Password         *CredentialsPassword         `json:"password"`
	OauthClient      CredentialsOauthClient       `json:"oauthClient"`
}

// See https://developer.okta.com/docs/api/resources/apps.html#username-template-object
type CredentialsUserNameTemplate *struct {
	// NONE, BUILT_IN, CUSTOM
	Type string `json:"type"`

	Template   *string `json:"template"`
	UserSuffix string  `json:"userSuffix"`
}

// See https://developer.okta.com/docs/api/resources/apps.html#signing-credential-object
type CredentialsSigning struct {
	Kid string `json:"kid"`
}

type CredentialsPassword struct {
	// This will only be used to create/update
	// It is not returned in GETs
	Value *string `json:"password"`
}

// See https://developer.okta.com/docs/api/resources/apps.html#oauth-credential-object
type CredentialsOauthClient struct {
	ClientID                *string `json:"client_id"`
	ClientSecret            *string `json:"client_secret"`
	TokenEndpointAuthMethod string  `json:"token_endpoint_auth_method"`
	AutoKeyRotation         *bool   `json:"autoKeyRotation"`
}
