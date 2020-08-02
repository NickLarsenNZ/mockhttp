# `mockhttp`

A testing utility for serving static responses.

## How to use

### Standalone

You can run a containerised instance

```sh
alias mockhttp='docker run --rm --network=host -u $(id -u) -v $(pwd):$(pwd):ro -w $(pwd) nicklarsennz/mockhttp'
mockhttp list  # Get a tabular list of responders listed in responders.yml
mockhttp serve # Serve the responders in responders.yml on http://localhost:8080
```

<details>
 <summary>Example responders.yml</summary>

```yml
responders:
- when:
    http:
      method: GET
      path: /things
  then:
    http:
      status: 200
    headers:
      Content-Type: application/json
    body: |
      [
          {"Name": "thing1"},
          {"Name": "thing2"},
          {"Name": "thing3"},
      ]
```
</details>

### As part of your go tests

<details>
 <summary>Assuming some code to be tested</summary>

```go
package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	baseUrl string
	client  *http.Client
}

type Thing struct {
	Name string
}

func (app *App) GetListOfThings() ([]Thing, error) {
	res, err := app.client.Get(app.baseUrl + "/things")
	things := make([]Thing)
	err := json.NewDecoder(res.Body).Decode(&things)
	return things, err
}

func NewApp(baseUrl string, client *http.Client) *App {
	return &App{
		baseUrl: baseUrl,
		client:  client,
	}
}
```
</details>

Create a YAML file defining the HTTP responses to return, given conditions:

```yml
responders:
- when:
    http:
      method: GET
      path: /things
  then:
    http:
      status: 200
    headers:
      Content-Type: application/json
    body: |
      [
          {"Name": "thing1"},
          {"Name": "thing2"},
          {"Name": "thing3"},
      ]
```

Instantiate the mock `http.Client` and hand it to your app to use. `mockhttp` will catch the requests, match them to the `when` conditions and return the `then` response.

```go
package main

import (
	"testing"

	mockhttp "github.com/nicklarsennz/mockhttp"
)

func TestSomeThing(t *testing.T) {
	// Instantiate a new http.Client
	client, err := mockhttp.NewClient("./fakes.yml")
	if err != nil {
		t.Errorf(errors.Wrap(err, "mockhttp.NewClient()").Error())
	}

	// Inject the mock client into the real app
	app := NewApp("http://localhost/8080", client)
	list, err := app.GetListOfThings()
	if err != nil {
		t.Errorf(errors.Wrap(err, "app.GetListOfThings()").Error())
	}

	expected := 3
	actual := len(list)
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
```
