package api

import (
	"net/http"
	"testing"
)

func TestClient_Auth(t *testing.T) {
	session := NewMockSession()
	session.Handle(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authorization := r.Header.Get("Authorization"); authorization != "SSWS mockapitoken" {
			t.Errorf("expected authorization header, got %s", authorization)
			http.Error(w, `{"errorCode":"E0000005","errorSummary":"Invalid session","errorLink":"E0000005","errorId":"oaez3UPmz2vS7OKQazqTzS1TQ","errorCauses":[]}`, http.StatusForbidden)
			return
		}
		w.Write([]byte(`[]`))
	}))
	client := NewClient(session)

	// This test uses the apps List() function to verify authorization
	// We don't care about the results; just whether an error is returned
	client.Apps().List()
}
