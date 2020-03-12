# auth-service/Dockerfile

# Build the auth-service
FROM golang:alpine AS build

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /go/src/app
WORKDIR /go/src/app

ENV GO111MODULE=on

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-service

# Copy the newly built auth-service to Alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=build /go/src/app/auth-service .