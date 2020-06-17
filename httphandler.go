package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type config struct {
	envEnvironment string
	envPort        string
	envAppName     string
	envDebugMode   bool
}

func loggingHandler(next http.HandlerFunc, cfg config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// switch out response writer for a recorder
		// for all subsequent handlers
		c := httptest.NewRecorder()
		next(c, r)

		if cfg.envDebugMode {
			log.Println(fmt.Sprintf("%q", c.Body))
		}

		// copy everything from response recorder
		// to actual response writer
		for k, v := range c.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(c.Code)
		c.Body.WriteTo(w)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().Format(time.RFC1123))
}
