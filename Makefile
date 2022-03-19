go build: go-http-server
all: clean dependencies build

re: clean test build


clean:
	@echo "Cleanning..."
	-rm -f $(BINARY_NAME)
	-rm -f $(BINARY_UNIX)
	-find . -type d -name -exec rm -rf \{} +
	-$(GOCLEAN) -i
	@echo "Done cleanning."

dependencies:
	@echo "Getting dependencies..."
	$(GOMOD) download
	$(GOMOD) vendor
	@echo "Done getting dependencies."

go-http-server:
	@echo "Building go http server"
# 	@echo "Installing vendors..."
# 	go install ./vendor/...
	@echo "Building..."
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v

serve:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME) serve

update-dependencies: dependencies
	$(GOGET) -u

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v


GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOLIST=$(GOCMD) list
GOTEST=$(GOCMD) test
GOGET=GO111MODULE=on $(GOCMD) get
GOMOD=GO111MODULE=on $(GOCMD) mod
BINARY_NAME=go-http-server
BINARY_UNIX=$(BINARY_NAME)_unix
