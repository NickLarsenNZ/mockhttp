package mockhttp

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/nicklarsennz/mockhttp/responders"
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

	// Check HTTP Status
	expectedStatus := 200
	actualStatus := response.StatusCode
	if expectedStatus != actualStatus {
		t.Errorf("\nexpected Status: '%d'\ngot Status: '%d', %s", expectedStatus, actualStatus, response.Status)
	}

	// Check Headers
	expectedContentType := "text/plain"
	actualContentType := response.Header.Get("Content-Type")

	if expectedContentType != actualContentType {
		t.Errorf("\nexpected Content-Type: '%s'\ngot Content-Type: '%s'", expectedContentType, actualContentType)
	}

	// Check Body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(errors.Wrap(err, "ioutil.ReadAll").Error())
	}

	expectedBody := "Hello"
	actualBody := strings.TrimSpace(string(body))

	if expectedBody != actualBody {
		t.Errorf("\nexpected body: '%s'\nactual body: '%s'", expectedBody, actualBody)
	}
}

func TestMatchResponseQueryParam(t *testing.T) {
	req, err := http.NewRequest("GET", "http://anything/api?offset=4", nil)
	if err != nil {
		t.Error(errors.Wrap(err, "http.NewRequest").Error())
	}

	config := &responders.ResponderConfig{
		Responders: []responders.Responder{
			{
				When: responders.WhenRequest{
					Http: responders.WhenHttp{
						Method: "GET",
						Path:   "/api",
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
			{
				When: responders.WhenRequest{
					Http: responders.WhenHttp{
						Method: "GET",
						Path:   "/api?offset=4",
					},
				},
				Then: responders.ThenResponse{
					Http: responders.ThenHttp{
						Status:  200,
						Message: "OK",
					},
					Headers: map[string]string{"Content-Type": "text/plain"},
					Body:    "Goodbye",
				},
			},
		},
	}

	var response *http.Response
	response = MatchResponse(req, config)
	if response == nil {
		t.Fatalf("nil response")
	}

	// Check Body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(errors.Wrap(err, "ioutil.ReadAll").Error())
	}

	expectedBody := "Goodbye"
	actualBody := strings.TrimSpace(string(body))

	if expectedBody != actualBody {
		t.Errorf("\nexpected body: '%s'\nactual body: '%s'", expectedBody, actualBody)
	}
}
