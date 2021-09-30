package main

import (
	"github.com/blacksfk/are_hub/http"
	"github.com/blacksfk/are_hub/http/middleware/auth"
	"github.com/blacksfk/are_hub/http/middleware/validate"
	"github.com/blacksfk/uf"
	"nhooyr.io/websocket"
)

// HTTP route definitions.
func routes(s *uf.Server, services *services) {
	a := auth.New(services.users)

	// channel routes
	c := http.NewChannel(services.channels)
	cv := validate.NewChannel()

	s.NewGroup("/channel").
		Get(c.Index).
		Post(c.Store, a.Request, cv.Store)

	s.NewGroup("/channel/:id").
		Get(c.Show, a.Request).
		Post(c.VerifyPassword, cv.Verify).
		Put(c.Update, a.Request, cv.Store).
		Delete(c.Delete, a.Request)

	// user routes
	u := http.NewUser(services.users)
	uv := validate.NewUser()

	s.Post("/login", u.Login, uv.Login)

	// websocket upgrade routes
	opts := &websocket.AcceptOptions{
		InsecureSkipVerify: true,
		CompressionMode:    websocket.CompressionDisabled,
	}

	ts := http.NewTelemetryServer(opts, services.channels)

	s.Get("/subscribe/:id", ts.Subscribe)
	// s.Get("/publish/:id", ts.Publish)
	s.Post("/publish/:id", ts.Publish)
}
