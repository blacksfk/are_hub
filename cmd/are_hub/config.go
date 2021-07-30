package main

import "flag"

type config struct {
	address            string
	allowOrigin        string
	dbUser             string
	dbPass             string
	dbHost             string
	dbName             string
	mongoAuthMechanism string
}

// Load configuration from the array of args containing flags and associated variables.
func load(args []string) (*config, error) {
	conf := &config{}
	set := flag.NewFlagSet("", flag.ExitOnError)

	// define flags
	// application/cors parameters
	set.StringVar(&conf.address, "address", ":6060",
		"Address for the server to listen on")
	set.StringVar(&conf.allowOrigin, "allow-origin", "*",
		"Access-Control-Allow-Origin value")

	// database parameters
	set.StringVar(&conf.dbUser, "db-user", "dev", "Database user")
	set.StringVar(&conf.dbPass, "db-pass", "dev", "Database user's password")
	set.StringVar(&conf.dbHost, "db-host", "localhost:27017", "Database host address")
	set.StringVar(&conf.dbName, "db-name", "are_hub",
		"Database name within the database server")

	// mongodb specific parameters
	set.StringVar(&conf.mongoAuthMechanism, "mongodb-auth-mechanism", "SCRAM-SHA-256",
		"MongoDB authentication mechanism")

	// parse the slice as flags
	return conf, set.Parse(args)
}
