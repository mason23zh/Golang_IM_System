package main

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	ServeIp    string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) (*Client, error) {
	client := &Client{
		ServeIp:    serverIp,
		ServerPort: serverPort,
	}

	// connect to server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		log.Println("net.Dial error:", err)
		return nil, err
	}

	client.conn = conn
	return client, nil
}

func main() {
	_, err := NewClient("127.0.0.1", 8080)
	if err != nil {
		fmt.Println(">>>Error connecting to server...")
	}

	fmt.Println(">>>Success connecting to server...l")
	select {}
}
