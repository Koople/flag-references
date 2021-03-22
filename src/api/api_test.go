package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_api_get_flags(t *testing.T) {
	apiKey := "test-api-key"
	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		requestApiKey := req.Header["X-Api-Key"][0]

		if requestApiKey != apiKey {
			t.Fail()
		}
		res.WriteHeader(200)
		fmt.Fprintln(res, `["flag1", "flag2"]`)
	}))
	defer ts.Close()

	options := KPLOptions{BaseUri: ts.URL, ApiKey: apiKey}
	sut := NewClient(options)

	flags, err := sut.GetListFlags()
	if err != nil {
		t.Errorf("Unexpected error on request: %s", err)
	}

	assert.Equal(t, flags, []string{"flag1", "flag2"})
}
