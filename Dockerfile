# Use the official Golang image as the base image
FROM golang:1.20-alpine3.18 AS build

LABEL name="FORUM PROJECT"
LABEL description="zone01 Dakar"
LABEL authors="mousdieng; abdouazindiaye; dgaye; pka; mnom"

RUN mkdir /app
RUN apk update && apk add bash && apk add tree
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary
RUN go build -o app .

# Use a minimal image as the final base image
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/app .
EXPOSE 8080
CMD ["./app"]
