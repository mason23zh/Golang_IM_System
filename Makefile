APP_NAME := im-system
BIN_DIR := bin

CLIENT_PKG := ./cmd/client
SERVER_PKG := ./cmd/server

LINUX_ARCH := amd64
WINDOWS_ARCH := amd64

.PHONY: all build linux windows client server clean run-server run-client

all: build

build: linux windows

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

linux: $(BIN_DIR)
	GOOS=linux GOARCH=$(LINUX_ARCH) go build -o $(BIN_DIR)/im-client-linux-$(LINUX_ARCH) $(CLIENT_PKG)
	GOOS=linux GOARCH=$(LINUX_ARCH) go build -o $(BIN_DIR)/im-server-linux-$(LINUX_ARCH) $(SERVER_PKG)

windows: $(BIN_DIR)
	GOOS=windows GOARCH=$(WINDOWS_ARCH) go build -o $(BIN_DIR)/im-client-windows-$(WINDOWS_ARCH).exe $(CLIENT_PKG)
	GOOS=windows GOARCH=$(WINDOWS_ARCH) go build -o $(BIN_DIR)/im-server-windows-$(WINDOWS_ARCH).exe $(SERVER_PKG)

client: $(BIN_DIR)
	go build -o $(BIN_DIR)/im-client $(CLIENT_PKG)

server: $(BIN_DIR)
	go build -o $(BIN_DIR)/im-server $(SERVER_PKG)

run-server:
	go run $(SERVER_PKG)

run-client:
	go run $(CLIENT_PKG)

clean:
	rm -rf $(BIN_DIR)
