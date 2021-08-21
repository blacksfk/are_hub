/*
Initial setup for the hub application. Seeds the database
with an administrator user defined in the command line
arguments.
*/
package main

import (
	"context"
	"flag"
	"log"

	"github.com/blacksfk/are_hub"
	"github.com/blacksfk/are_hub/hash"
	"github.com/blacksfk/are_hub/mongodb"
	"github.com/go-playground/validator/v10"
)

type user struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

type config struct {
	dbUser             string
	dbPass             string
	dbHost             string
	dbName             string
	mongoAuthMechanism string
	admin              user
}

func main() {
	conf := loadConfig()
	validate := validator.New()

	// validate the admin user
	e := validate.Struct(conf.admin)

	if e != nil {
		// validation failed
		log.Fatal(e)
	}

	// valid; connect to the mongodb instance
	params := mongodb.Params{
		User:      conf.dbUser,
		Password:  conf.dbPass,
		Mechanism: conf.mongoAuthMechanism,
		Address:   conf.dbHost,
		Name:      conf.dbName,
	}

	ctx := context.Background()
	client, e := mongodb.Connect(ctx, &params)

	if e != nil {
		// connection failed
		log.Fatal(e)
	}

	// hash the password
	hashed, e := hash.Password(conf.admin.Password)

	if e != nil {
		// password hashing failed
		log.Fatal(e)
	}

	// create the users collection and insert the admin user
	users := mongodb.NewUserCollection(client, conf.dbName)
	admin := are_hub.NewUser(conf.admin.Name, hashed)
	e = users.Insert(ctx, admin)

	if e != nil {
		// insertion failed
		log.Fatal(e)
	}

	log.Printf("Inserted: %+v\n", admin)
}

func loadConfig() config {
	conf := config{admin: user{}}

	// define flags for the admin user
	flag.StringVar(&conf.admin.Name, "name", "", "Admin username")
	flag.StringVar(&conf.admin.Password, "password", "", "Admin password")

	// define flags for the database
	flag.StringVar(&conf.dbUser, "db-user", "dev", "Database user")
	flag.StringVar(&conf.dbPass, "db-pass", "dev", "Database user's password")
	flag.StringVar(&conf.dbHost, "db-host", "localhost:27017", "Database host address")
	flag.StringVar(&conf.dbName, "db-name", "are_hub",
		"Database name within the database server")

	// mongodb specific parameters
	flag.StringVar(&conf.mongoAuthMechanism, "mongodb-auth-mechanism", "SCRAM-SHA-256",
		"MongoDB authentication mechanism")

	// parse the commandline args
	flag.Parse()

	return conf
}
