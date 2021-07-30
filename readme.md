# ACC Race Engineer hub

This is the hub application that receives data and forwards it to the appropriate connected websocket clients.

## Compiling and running
Your `go version` must support modules in order for `go build` to obtain the necessary dependencies. Currently `mongodb` is the only supported database.

1. Install mongodb and create a new database.
2. `cd cmd/are_hub/`
3. `go build`
4. `./are_hub` or `./are_hub --help` to view the commandline arguments and their default values.

## Deployment
Building a docker container is the easiest way (probably).

1. `docker build -t are_hub:<tag> .`
2. `docker run -d -p 9001:9001 --network backend are_hub:<tag> --address :9001 --allow-origin example.com --db-user blast_hardcheese --db-pass butch_deadlift --db-host mongodb --db-name acc_race_engineer`

## Licence
BSD-3-clause
