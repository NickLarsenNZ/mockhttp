package mockhttp

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/nicklarsennz/mock-http-response/responders"
	"github.com/pkg/errors"
)

func NewClient(definitionsFilePath string) (*http.Client, error) {
	config, err := responders.ParseConfig(definitionsFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "parser error")
	}

	client := &http.Client{
		Transport: &mockTransport{
			ResponderConfig: config,
		},
	}

	return client, nil
}

type mockTransport struct {
	*responders.ResponderConfig
}

func (c *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	//fmt.Println("Request URL:", req.URL)

	// Return a fake response without any interaction with the network
	response := &http.Response{
		StatusCode: 200,
		Header:     http.Header{"X-NOOP": []string{"blah"}},
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"success": true}`)),
	}

	return response, nil
}
