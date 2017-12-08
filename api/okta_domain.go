package api

import "fmt"

type OktaDomain struct {
	Secure  bool
	BaseURL string
	OrgName string
}

func (d OktaDomain) String() string {
	return fmt.Sprintf("%s://%s.%s", d.Scheme(), d.OrgName, d.BaseURL)
}

func (d OktaDomain) Scheme() string {
	if d.Secure {
		return "https"
	}
	return "http"
}
