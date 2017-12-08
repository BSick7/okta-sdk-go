package api

/// All implementations of Session are expected to use the following environment variables:
///   OKTA_ORG_NAME
///   OKTA_API_TOKEN
///   OKTA_DEBUG (turns on request/response logging)

import (
	"fmt"
	"net/http"
	"net/url"
)

type Session interface {
	SetDomain(domain fmt.Stringer)
	SetAPIToken(apiToken string)
	EndpointURL(endpoint string) *url.URL
	CreateTransport() http.RoundTripper
}
