# Use Golang as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy all the source code into the container
COPY . .

# Expose port 8080 for the web server
EXPOSE 8080

# Run the application
CMD ["go", "run", "cmd/main.go"]
