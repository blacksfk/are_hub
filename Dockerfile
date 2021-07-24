# alpine distro
FROM golang:alpine AS builder

# create and change to application directory
WORKDIR /app/

# copy source files
COPY . ./

# change to build directory and compile
RUN go mod download
WORKDIR ./cmd/are_hub/
RUN go build

# separate image for running
FROM alpine

COPY --from=builder /app/cmd/are_hub/are_hub .
COPY config.json.docker ./config.json

# start the application
# remember to expose the port set in config.json!
ENTRYPOINT ["./are_hub", "--config", "./config.json"]
