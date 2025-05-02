# Use the official Golang image as the base
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire project directory
COPY . .

# Copy the .env file into the container
COPY .env .env

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 8081

# Set the default command to run the executable
CMD ["./main"]
