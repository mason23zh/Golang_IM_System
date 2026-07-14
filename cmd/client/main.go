package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

type Client struct {
	ServeIp    string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //current client mode
}

func (c *Client) menu() bool {
	// user input
	var flag int

	fmt.Println("1. Public Chat")
	fmt.Println("2. Private Chat")
	fmt.Println("3. Update user name")
	fmt.Println("0. Exit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println(">>>input out of range<<<")
		return false
	}
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

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8080
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "setting server IP address")
	flag.IntVar(&serverPort, "port", 8080, "setting server port number")
}

func main() {
	//cli parsing
	flag.Parse()

	_, err := NewClient(serverIp, serverPort)
	if err != nil {
		fmt.Println(">>>Error connecting to server...")
	}

	fmt.Println(">>>Success connecting to server...l")
	time.Sleep(5 * time.Minute)
}
