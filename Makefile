.PHONY: build install uninstall

BINARY := rubber-duck
BUILD_DIR := bin
BUILD_PATH := $(BUILD_DIR)/$(BINARY)
INSTALL_DIR := $(HOME)/bin
INSTALL_PATH := $(INSTALL_DIR)/$(BINARY)

build:
	go build -o $(BUILD_PATH) ./cmd/rubber-duck

install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BUILD_PATH) $(INSTALL_PATH)

uninstall:
	rm -f $(INSTALL_PATH)
