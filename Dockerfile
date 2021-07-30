# alpine distro
FROM golang:alpine AS builder

# create and change to application directory
WORKDIR /app/

# copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy source files
COPY . ./

# change to build directory and compile
WORKDIR ./cmd/are_hub/
RUN go build

# separate image for running
FROM alpine

WORKDIR /app/
RUN mkdir config

COPY --from=builder /app/cmd/are_hub/are_hub .

# start the application
ENTRYPOINT ["./are_hub"]
