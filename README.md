# `mock-http-response`

## How to use

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
    client *http.Client
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
        client: client,
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

Instantiate the mock `http.Client` and hand it to your app to use. `mock-http-responses` will catch the requests, match them to the `when` conditions and return the `then` response.

```go
package main

import (
    mockhttp "github.com/nicklarsennz/mock-http-responder"
)

func TestSomeThing(t *testing.T) {
    // Instantiate a new http.Client
    client := mockhttp.NewClient("./fakes.yml")

    // Inject the mock client into the real app
    app := NewApp(client, "http://localhost/8080")
    list, err := app.GetListOfThings()
    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }

    expected := 3
    actual := len(list)
    if actual != expected {
        t.Errorf("expected %d, got %d", expected, actual)
    }
}
```
