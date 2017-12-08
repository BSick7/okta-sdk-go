package api

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	DEFAULT_BASE_URL = "okta.com"
	PREVIEW_BASE_URL = "oktapreview.com"
)

type SessionImpl struct {
	domain   fmt.Stringer
	apiToken string
}

func DefaultSession() Session {
	s := &SessionImpl{}
	baseURL := os.Getenv("OKTA_DOMAIN")
	if baseURL == "" {
		baseURL = DEFAULT_BASE_URL
	}
	s.SetDomain(OktaDomain{
		Secure:  true,
		BaseURL: baseURL,
		OrgName: os.Getenv("OKTA_ORG_NAME"),
	})
	s.SetAPIToken(os.Getenv("OKTA_API_TOKEN"))
	return s
}

func (s *SessionImpl) SetDomain(domain fmt.Stringer) {
	s.domain = domain
}

func (s *SessionImpl) SetAPIToken(apiToken string) {
	s.apiToken = apiToken
}

func (s *SessionImpl) Domain() string {
	return s.domain.String()
}

func (s *SessionImpl) EndpointURL(endpoint string) *url.URL {
	uri, _ := url.Parse(fmt.Sprintf("%s/api/v1/%s", strings.TrimRight(s.Domain(), "/"), strings.TrimLeft(endpoint, "/")))
	return uri
}

func (s *SessionImpl) CreateTransport() http.RoundTripper {
	return NewAnonymousTransport(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("User-Agent", "okta-sdk-go/0.1.0")
		req.Header.Set("Authorization", fmt.Sprintf("SSWS %s", s.apiToken))
		return http.DefaultTransport.RoundTrip(req)
	})
}
