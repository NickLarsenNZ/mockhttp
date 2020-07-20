package responders

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Headers map[string]string

func (h Headers) AppearIn(headers http.Header) bool {
	// Default to true, as no headers mean allow anything
	var matches = true
	// For each of our required headers
	for k, v := range h {
		// If we have at least one header, then it better match
		matches = false
		// http.Header has a list of values
		for _, value := range headers.Values(k) {
			if v == value {
				// Cool, move on to the next header
				break
			}
		}
		// If there were no matches, return false straight away
		if !matches {
			return matches
		}
		// Otherwise, move on to the next header
	}
	return matches // this will always be true
}

type WhenHttp struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}

type ThenHttp struct {
	Status  int    `yaml:"status"`
	Message string `message:"message"`
}

// Represents the "when" block under a response object
type WhenRequest struct {
	Http    WhenHttp `yaml:"http"`
	Headers Headers  `yaml:"headers"`
	Body    string   `yaml:"body,omitempty"`
}

// Represents the "then" block under a response object
type ThenResponse struct {
	Http    ThenHttp          `yaml:"http"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}

type Responder struct {
	When WhenRequest  `yaml:"when"`
	Then ThenResponse `yaml:"then"`
}

type ResponderConfig struct {
	Responders []Responder `yaml:"responders"`
}

// Parse the configuration for a set of responders
// Given the following YAML:
//  ---
//  responders:
//  # Initial page
//  - when:
//      http:
//        method: GET
//        path: /
//    then:
//      http:
//        status: 200
//      headers:
//        Content-Type: text/html
//      body: |
//        <html>
//          <body>
//            <form method="post" action="/login">
//              <input type="submit" value="Login" />
//            </form>
//          </body>
//        </html>
// You can load that:
//  config, err := responders.ParseConfig("/path/to/spec.yaml")
//  if err != nil {
//    panic(err.Error())
//  }
//  responder_count = len(config.Responders) // 1
func ParseConfig(yamlFilePath string) (*ResponderConfig, error) {

	var r *ResponderConfig

	yamlFile, err := os.Open(yamlFilePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = yamlFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	yamlBytes, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlBytes, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

var defaultResponseMessages = map[int]string{
	100: "Continue",
	101: "Switching Protocols",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	307: "Temporary Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Time-out",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Request Entity Too Large",
	414: "Request-URI Too Large",
	415: "Unsupported Media Type",
	416: "Requested range not satisfiable",
	417: "Expectation Failed",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Time-out",
	505: "HTTP Version not supported",
}
