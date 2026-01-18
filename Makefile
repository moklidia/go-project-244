build:
	go build -o gendiff cmd/gendiff/main.go

install:
	go install -o gendiff cmd/gendiff/main.go

run:
	go run cmd/gendiff/main.go

test:
	go test ./...

lint:
	golangci-lint run ./...
