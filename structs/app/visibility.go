package app

// See https://developer.okta.com/docs/api/resources/apps.html#visibility-object
type Visibility struct {
	AutoSubmitToolbar bool `json:"autoSubmitToolbar"`
	Hide              struct {
		IOS bool `json:"iOS"`
		Web bool `json:"web"`
	} `json:"hide"`
}
