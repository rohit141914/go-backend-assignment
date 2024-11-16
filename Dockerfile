# Use a lightweight Alpine image as base
FROM alpine:latest

# Install dependencies
RUN apk add --no-cache \
    bash \
    curl \
    && curl -LO https://golang.org/dl/go1.22.3.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz \
    && rm go1.22.3.linux-amd64.tar.gz

# Set Go binary to the PATH
ENV PATH="/usr/local/go/bin:${PATH}"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["go", "run", "cmd/main.go"]
