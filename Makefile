
APP_NAME := im-system
BIN_DIR := bin

CLIENT_PKG := ./cmd/client
SERVER_PKG := ./cmd/server

LINUX_ARCH := amd64
WINDOWS_ARCH := amd64

.PHONY: all build linux windows client server clean run-server run-client verify-windows

all: build

build: linux windows

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

linux: $(BIN_DIR)
	env GOOS=linux GOARCH=$(LINUX_ARCH) CGO_ENABLED=0 go build -o $(BIN_DIR)/im-client-linux-$(LINUX_ARCH) $(CLIENT_PKG)
	env GOOS=linux GOARCH=$(LINUX_ARCH) CGO_ENABLED=0 go build -o $(BIN_DIR)/im-server-linux-$(LINUX_ARCH) $(SERVER_PKG)

windows: $(BIN_DIR)
	env GOOS=windows GOARCH=$(WINDOWS_ARCH) CGO_ENABLED=0 go build -o $(BIN_DIR)/im-client-windows-$(WINDOWS_ARCH).exe $(CLIENT_PKG)
	env GOOS=windows GOARCH=$(WINDOWS_ARCH) CGO_ENABLED=0 go build -o $(BIN_DIR)/im-server-windows-$(WINDOWS_ARCH).exe $(SERVER_PKG)

client: $(BIN_DIR)
	go build -o $(BIN_DIR)/im-client $(CLIENT_PKG)

server: $(BIN_DIR)
	go build -o $(BIN_DIR)/im-server $(SERVER_PKG)

run-server:
	go run $(SERVER_PKG)

run-client:
	go run $(CLIENT_PKG)

verify-windows:
	file $(BIN_DIR)/im-client-windows-$(WINDOWS_ARCH).exe
	file $(BIN_DIR)/im-server-windows-$(WINDOWS_ARCH).exe

clean:
	rm -rf $(BIN_DIR)
