GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/iss-notifier
BINARY_UNIX=$(BINARY_NAME)_unix

setup:
	go get -u github.com/kardianos/govendor

vendoring:
	govendor remove +external
	govendor add +external
	govendor update +external

build-lambda:
	@echo "Building ISS Notifier for AWS Lambda"
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o main main.go
	@zip main.zip main
	mv main.zip dist/main.zip
	rm main

test:
	@echo "Running iss-notifier tests"
	$(GOTEST) -v ./...

fmt:
	@echo "Running gofmt for all project files"
	$(GOFMT) -w *.go