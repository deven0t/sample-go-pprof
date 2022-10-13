build:
	GOOS=linux GOARCH=amd64 go build main.go

lint:
	golangci-lint run -v ./...

image:
	docker build -t turkardg/sample-go-pprof:linux .