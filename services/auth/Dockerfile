# Dockerfile for each Go service
FROM golang:1.23.2

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
# COPY go.sum ./
RUN go mod download

# Copy the service code
COPY . .

# Build the application
RUN go build -o auth .

# Expose port (change as needed for each service)
EXPOSE 8081

# Run the executable
CMD ["./auth"]
