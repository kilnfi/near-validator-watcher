package testutils

import (
	"net/http"
	"net/http/httptest"
)

type ExpectedResponse struct {
	StatusCode int
	Body       []byte
}

func (r *ExpectedResponse) ExpectResponse(statusCode int, body string) {
	r.StatusCode = statusCode
	r.Body = []byte(body)
}

func NewServer(resp *ExpectedResponse) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(resp.StatusCode)
		res.Write(resp.Body)
	}))
}
