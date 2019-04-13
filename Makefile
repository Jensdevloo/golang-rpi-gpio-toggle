GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=golang-rpi-gpio-toggle
BINARY_UNIX=golang-rpi-gpio-toggle-arm

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

dep:
	dep ensure

build-rpi:
	env GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -o $(BINARY_UNIX) -v

upload-to-pi:
	scp golang-rpi-gpio-toggle-arm pi@dragonsdenpi.local:~

docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
