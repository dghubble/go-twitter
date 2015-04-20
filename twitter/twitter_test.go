package twitter

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

// testing utils

// testServer returns an http Client, ServeMux, and Server. The client proxies
// requests to the server and handlers can be registered on the mux to handle
// requests. The caller must close the test server.
func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}
	client := &http.Client{Transport: transport}
	return client, mux, server
}

// RewriteTransport rewrites https requests to http to avoid TLS cert issues
// during testing.
type RewriteTransport struct {
	Transport http.RoundTripper
}

// RoundTrip rewrites the request scheme to http and calls through to the
// composed RoundTripper or if it is nil, to the http.DefaultTransport.
func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func assertMethod(t *testing.T, expectedMethod string, req *http.Request) {
	if actualMethod := req.Method; actualMethod != expectedMethod {
		t.Errorf("expected method %s, got %s", expectedMethod, actualMethod)
	}
}

func assertParams(t *testing.T, expected map[string]string, req *http.Request) {
	req.ParseForm() // populates r.Form url.values (alias map[string][]string)
	expectedValues := url.Values{}
	for key, value := range expected {
		expectedValues.Add(key, value)
	}
	if !reflect.DeepEqual(expectedValues, req.Form) {
		t.Errorf("expected parameters %v, got %v", expected, req.Form)
	}
}
