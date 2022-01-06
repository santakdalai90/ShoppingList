GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOLINT=golint
BINARY_NAME=shopping_list

.PHONY: test

all: test clean build

## lint: Runs the linter on the source
lint:
	$(GOLINT) -set_exit_status ./...

## build: Gathers the dependencies and builds the binary
build:
	$(GOMOD) tidy
	$(GOBUILD) -o $(BINARY_NAME) -v

## test: Runs all available tests in the source code
test:
	$(GOTEST) -v ./...

## clean: Cleans the build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

## run: Builds and runs the project
run: build
	./$(BINARY_NAME)

## help: gives help instructions
help: Makefile
	@echo
	@echo "Available Commands:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo