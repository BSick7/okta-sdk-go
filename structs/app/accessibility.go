package app

// See https://developer.okta.com/docs/api/resources/apps.html#accessibility-object
type Accessibility struct {
	SelfService      *bool  `json:"selfService"`
	ErrorRedirectUrl string `json:"errorRedirectUrl"`
	LoginRedirectUrl string `json:"loginRedirectUrl"`
}
