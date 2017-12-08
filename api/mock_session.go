package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type MockSession struct {
	impl    *SessionImpl
	server  *httptest.Server
	handler http.HandlerFunc
}

type mockDomain struct{ server *httptest.Server }

func (d mockDomain) String() string { return d.server.URL }

func NewMockSession() *MockSession {
	s := &MockSession{
		handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}

	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.handler(w, r)
	}))

	s.impl = &SessionImpl{
		domain:   mockDomain{s.server},
		apiToken: "mockapitoken",
	}

	return s
}

func (s *MockSession) Handle(handler http.HandlerFunc) {
	s.handler = handler
}

func (s *MockSession) SetDomain(domain fmt.Stringer) {
	// no-op
	// domain should hit http test server
}

func (s *MockSession) SetAPIToken(apiToken string) {
	s.impl.SetAPIToken(apiToken)
}

func (s *MockSession) EndpointURL(endpoint string) *url.URL {
	return s.impl.EndpointURL(endpoint)
}

func (s *MockSession) CreateTransport() http.RoundTripper {
	return s.impl.CreateTransport()
}
