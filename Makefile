.PHONY: all build run clean

# Name of the binary output file
BINARY_NAME := muxapi

all: build

build:
	go build -o $(BINARY_NAME) main.go

run:
	go run main.go

clean:
	rm -f $(BINARY_NAME)
