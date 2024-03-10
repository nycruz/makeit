build: # Builds the project 
	go build -o makeit

run: # Build/Run the project
	go run *.go -f "Makefile.example"

clean: # Clean the project
	rm -f makeit

test: # Run the tests
	go test -v ./...

.PHONY: build run
