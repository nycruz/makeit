build: # 🚧 Builds the project 
	go build -o bin/makeit

run: # 🏎️ Run the project
	go run *.go

.PHONY: build run
