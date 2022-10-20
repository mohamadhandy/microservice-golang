package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}
	log.Printf("Broker service run on port %s\n", webPort)

	// define http server
	serve := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
