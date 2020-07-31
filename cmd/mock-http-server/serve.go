package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	mockhttp "github.com/nicklarsennz/mock-http-response"
	"github.com/nicklarsennz/mock-http-response/responders"
)

func newServer(config *responders.ResponderConfig, bind_addr string, bind_port int) *http.Server {
	r := mux.NewRouter()

	for _, res_iter := range config.Responders {
		var responder = res_iter // https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		handler := func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Hit on:", r.Method, r.URL.Path)

			res := mockhttp.MatchResponse(r, config)

			for k, values := range res.Header {
				for _, v := range values {
					fmt.Printf("Setting: %s: %s\n", k, v)
					w.Header().Set(k, v)
				}
			}

			w.WriteHeader(res.StatusCode)

			io.Copy(w, res.Body) // https://stackoverflow.com/a/28891552
			res.Body.Close()
		}

		// Add, or update the handler for the given path
		r.HandleFunc(responder.When.Http.Path, handler).Methods(responder.When.Http.Method)
	}

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		m, _ := route.GetMethods()
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t, m)
		return nil
	})

	return &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%d", bind_addr, bind_port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
