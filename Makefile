GOCMD := go
BINARY := go-stock-ticker
BUILD_DIR := bin
SRC_DIR := src

.PHONY: clean

all: clean dep test build

dep:
	$(GOCMD) get github.com/olekukonko/tablewriter

clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

build:
	mkdir -p $(BUILD_DIR)
	cd $(SRC_DIR)/; $(GOCMD) build -o ../$(BUILD_DIR)/$(BINARY) 

test:
	cd $(SRC_DIR); $(GOCMD) test -v
run:
	$(BUILD_DIR)/$(BINARY) TSLA