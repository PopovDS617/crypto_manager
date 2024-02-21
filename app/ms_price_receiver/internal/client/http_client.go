package client

import (
	"time"

	"github.com/gojek/heimdall/httpclient"
)

type HTTPClient struct {
	Client *httpclient.Client
	Url    string
}

func NewHTTPClient(url string) *HTTPClient {

	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	return &HTTPClient{Client: client, Url: url}
}
