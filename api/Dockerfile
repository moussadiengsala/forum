# Use the official Go image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o main cmd/main.go

# Expose the port the app runs on
EXPOSE 7000

# Command to run the executable
CMD ["./main"]

