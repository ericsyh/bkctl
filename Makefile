GOCMD := GO111MODULE=on go
GOBUILD := $(GOCMD) build
BINARY_NAME="bin/bkctl"

default: build

build:
	$(GOBUILD) -o $(BINARY_NAME) 