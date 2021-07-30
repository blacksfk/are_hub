package main

import (
	"log"
	"net/http"
	"os"

	uf "github.com/blacksfk/microframework"
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
		Address:      conf.address,
		ErrorLogger:  logStdout,
		AccessLogger: uf.LogStdout,
	}

	// create server
	s := uf.NewServer(sconf)

	// setup cors
	s.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// cors pre-flight request; reply with cors headers
			h := w.Header()

			// TODO: verify these are the only cors headers we need.
			// Possibly "Access-Control-Allow-Credentials" amongst others.
			h.Set("Access-Control-Allow-Methods", h.Get("Allow"))
			h.Set("Access-Control-Allow-Origin", conf.allowOrigin)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	// define routes
	routes(s, services)

	// anchors aweigh!
	log.Fatal(s.Start())
}

func logStdout(e error) {
	log.Println(e)
}
