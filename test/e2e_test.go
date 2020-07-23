package test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	mockhttp "github.com/nicklarsennz/mock-http-response"
	"github.com/pkg/errors"
)

func TestUtility(t *testing.T) {
	var expectedHeaders = map[string]string{
		"Content-Type": "application/json",
	}

	client, err := mockhttp.NewClient("./testdata/fakes.yml")
	if err != nil {
		t.Errorf(errors.Wrap(err, "mockhttp.NewClient").Error())
	}

	req, err := http.NewRequest("GET", "http://anything/things", nil)
	if err != nil {
		t.Errorf(errors.Wrap(err, "http.NewRequest").Error())
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf(errors.Wrap(err, "client.Do").Error())
	}
	defer res.Body.Close()

	for key, expectedValues := range expectedHeaders {
		// Ensure it is in the actual headers
		actualValues := strings.Join(req.Header.Values(key), ",")
		if expectedValues != actualValues {
			t.Errorf("\nexpected header (%s): '%s'\ngot header (%s): '%s'", key, expectedValues, key, actualValues)
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf(errors.Wrap(err, "ioutil.ReadAll").Error())
	}

	var expectedBody = strings.TrimSpace(`
[
    {"Name": "thing1"},
    {"Name": "thing2"},
    {"Name": "thing3"},
]`)
	actualBody := strings.TrimSpace(string(body))
	if expectedBody != actualBody {
		t.Errorf("\nexpected body: '%s'\nactual body: '%s'", expectedBody, actualBody)
	}
}
