# Define variables (adjust as needed)
GOARCH := amd64
GOOS := linux  # You can add other OS options here (e.g., windows, darwin)
BINARY_NAME := thegamem

build: ./cmd/thegamem.go  # Specify the dependency here
	go build -o $(BINARY_NAME) ./cmd/thegamem.go

clean:
	rm -f $(BINARY_NAME)

run: build  # Specify build as a dependency here
	./$(BINARY_NAME)