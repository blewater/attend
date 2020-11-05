run:
	go run srv.go
build:
	go build ./...

test:
	go test ./...

lint:
	golangci-lint run

# go get mvdan.cc/gofumpt
fmt:
	gofumports -w **/*.go

# go get github.com/daixiang0/gci
imp:
	gci -w **/*.go
