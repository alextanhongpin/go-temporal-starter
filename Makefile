start:
	go run main.go

lint:
	golangci-lint run

install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1

dashboard:
	@open localhost:8088
