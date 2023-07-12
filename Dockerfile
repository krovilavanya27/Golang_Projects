# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache the Go module dependencies
RUN go mod download

# Copy the source code and the Makefile
COPY . .

# Build the Go program using the Makefile
RUN make build

# Set the entry point to run the Go program
CMD ["make", "run"]
