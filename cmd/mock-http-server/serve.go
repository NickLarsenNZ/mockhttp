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

	for _, res := range config.Responders {
		var responder = res // https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		handler := func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Hit on:", r.Method, r.URL.Path)

			for k, v := range responder.Then.Headers {
				w.Header().Set(k, v)
			}

			w.WriteHeader(responder.Then.Http.Status)

			res := mockhttp.MatchResponse(r, config)
			io.Copy(w, res.Body) // https://stackoverflow.com/a/28891552
			res.Body.Close()
		}

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
