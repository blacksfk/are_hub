package main

import (
	"context"
	"log"

	"github.com/blacksfk/are_hub"
	"github.com/blacksfk/are_hub/mongodb"
)

// Wraps various database tables and services that
// require initilisation.
type services struct {
	channels are_hub.ChannelRepo
	users    are_hub.UserRepo
}

// Initialise various services and create mongodb collections based on conf. This function
// is only intended to be called from the main function therefore dies if it encounters
// an error creating a mongo.Client.
func initServices(conf *config) *services {
	params := mongodb.Params{
		User:      conf.dbUser,
		Password:  conf.dbPass,
		Mechanism: conf.mongoAuthMechanism,
		Address:   conf.dbHost,
		Name:      conf.dbName,
	}

	client, e := mongodb.Connect(context.Background(), &params)

	if e != nil {
		// connection failed so die
		log.Fatal(e)
	}

	return &services{
		mongodb.NewChannelCollection(client, conf.dbName),
		mongodb.NewUserCollection(client, conf.dbName),
	}
}
