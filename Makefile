GOCMD := /usr/bin/go
BINARY := go-stock-ticker
BUILD_DIR := build


all: test build

clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)


build:
	mkdir -p build
	$(GOCMD) build -o $(BUILD_DIR)/$(BINARY) 

test:
	$(GOCMD) test