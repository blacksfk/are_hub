package main

import (
	"log"
	"net/http"
	"os"

	uf "github.com/blacksfk/microframework"
	"github.com/rs/cors"
)

func main() {
	// load the configuration
	conf, e := load(os.Args[1:])

	if e != nil {
		log.Fatal(e)
	}

	// initialise services
	services := initServices(conf)

	// server configuration
	sconf := &uf.Config{
		ErrorLogger:  logStdout,
		AccessLogger: uf.LogStdout,
	}

	// create server
	s := uf.NewServer(sconf)

	// setup cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{conf.allowOrigin},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut,
			http.MethodPatch, http.MethodDelete},
	})

	// define routes
	routes(s, services)

	// anchors aweigh!
	log.Fatal(http.ListenAndServe(conf.address, c.Handler(s)))
}

func logStdout(e error) {
	log.Println(e)
}
