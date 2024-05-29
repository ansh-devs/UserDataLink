# Build the application
all: build

build:
	@echo "Building..."
	
	@go build -o main cmd/main.go

proto-gen:
	@  echo "[INFO]: generating protofiles. "; protoc --go_out=. \
  		--go-grpc_out=. protos/user_service.proto


# Run the application
run:
	@go run cmd/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

.PHONY: all build run test clean
