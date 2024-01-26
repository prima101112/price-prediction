# Use an official Golang runtime as a parent image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app
COPY ./migrations/seed/historicalpo.csv /app/migrations/seed/historicalpo.csv

# Build the Go application
RUN go build -o main ./cmd/server.go

# Expose port 8080 to the outside world
EXPOSE 8090

# Command to run the executable
CMD ["./main"]