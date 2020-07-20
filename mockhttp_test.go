package mockhttp

import (
	"net/http"
	"testing"

	"github.com/nicklarsennz/mock-http-response/responders"
	"github.com/pkg/errors"
)

func TestMatchResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "http://anything/blah", nil)
	if err != nil {
		t.Error(errors.Wrap(err, "http.NewRequest").Error())
	}

	config := &responders.ResponderConfig{
		Responders: []responders.Responder{
			{
				When: responders.WhenRequest{
					Http: responders.WhenHttp{
						Method: "GET",
						Path:   "/blah",
					},
				},
				Then: responders.ThenResponse{
					Http: responders.ThenHttp{
						Status:  200,
						Message: "OK",
					},
					Headers: map[string]string{"Content-Type": "text/plain"},
					Body:    "Hello",
				},
			},
		},
	}

	var response *http.Response
	response = MatchResponse(req, config)
	if response == nil {
		t.Fatalf("nil response")
	}

	expectedContentType := "text/plain"
	actualContentType := response.Header.Get("Content-Type")

	if expectedContentType != actualContentType {
		t.Errorf("\nexpected Content-Type: '%s'\ngot Content-Type: '%s'", expectedContentType, actualContentType)
	}
}
