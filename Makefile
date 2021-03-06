lint:
	golint ./...
test:
	go test -v -failfast -covermode=count ./...
coverage:
	go test ./... -cover
build:
	go build -o dietify
run:
	go run cmd/main.go
mod:
	go mod vendor
