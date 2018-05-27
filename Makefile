GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/iss-tracker
BINARY_UNIX=$(BINARY_NAME)_unix

setup:
	go get -u github.com/kardianos/govendor

vendoring:
	govendor remove +external
	govendor add +external
	govendor update +external

build-lambda:
	@echo "Building ISS Tracker for AWS Lambda"
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o main main.go
	@zip main.zip main
	mv main.zip dist/main.zip
	rm main

run-local:
	@echo "Running ISS Tracker in local mode"
	$(GOCMD) run main.go local

test:
	@echo "Running iss-tracker tests"
	$(GOTEST) -v ./...

fmt:
	@echo "Running gofmt for all project files"
	$(GOFMT) -w *.go