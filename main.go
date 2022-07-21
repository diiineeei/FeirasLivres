package main

import (
	"net/http"
	"os"
	"testesrod/handlers"
	"testesrod/logs"
	"time"
)

const (
	defaultPort       = "8888"
	idleTimeout       = 30 * time.Second
	writeTimeout      = 180 * time.Second
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 10 * time.Second
)

func main() {
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = defaultPort
	}

	server := &http.Server{
		Addr:              "0.0.0.0:" + Port,
		Handler:           handlers.New(),
		IdleTimeout:       idleTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		logs.Print("server.ListenAndServe %v ", err)
	}
}
