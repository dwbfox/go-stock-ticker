GOCMD := /usr/bin/go
BINARY := go-stock-ticker
BUILD_DIR := build
SRC_DIR := src/

.PHONY: clean build

all: clean test build

clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)


build:
	mkdir -p build
	$(GOCMD) build $(SRC_DIR)* -o $(BUILD_DIR)/$(BINARY) 

test:
	$(GOCMD) test $(SRC_DIR)/

run:
	$(BUILD_DIR)/$(BINARY) TSLA