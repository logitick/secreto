SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := "$(shell echo $${PWD\#\#*/})-$(shell uname -s | tr '[:upper:]' '[:lower:]')-$(shell uname -i)"
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"


# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")


.PHONY: all build docker clean install uninstall fmt simplify check run

#all: check install
all: deps test release

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

test: deps
	@go test ./...

docker: build
	@docker build -t secreto:latest --build-arg bin="$(TARGET)" .
	$(MAKE) clean

clean:
	@rm -rf release
	@rm -f $(TARGET)

install:
	@go install $(LDFLAGS)

deps:
	@go get

release-deps:
	@go get \
	github.com/mitchellh/gox \
	github.com/inconshreveable/mousetrap # for windows build

release: release-deps
	gox -verbose \
	-osarch="windows/amd64 linux/amd64 darwin/amd64" \
	-ldflags "-X main.version=${VERSION}" \
-output="release/{{.Dir}}-${VERSION}-{{.OS}}-{{.Arch}}" .

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

run: install
	@$(TARGET)