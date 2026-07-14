package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type Client struct {
	ServeIp    string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //current client mode
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8080
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "setting server IP address")
	flag.IntVar(&serverPort, "port", 8080, "setting server port number")
}

func (c *Client) Run() {
	for c.flag != 0 {
		for c.menu() != true {
		}
		switch c.flag {
		case 1:
			// public chat mode
			fmt.Println("public chat mode")
		case 2:
			// private chat mode
			fmt.Println("private chat mode")
		case 3:
			// change name mode
			c.UpdateName()
		}
	}
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
		flag:       -1, // cannot use zero value otherwise it will exit
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

// handle server response, display in stdout
func (c *Client) DealResponse() {
	// block here, once c.conn has data, copy the data into stdout
	io.Copy(os.Stdout, c.conn)
}

func (c *Client) UpdateName() bool {
	fmt.Println(">>>enter user name<<<")
	fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"
	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	return true
}

func main() {
	//cli parsing
	flag.Parse()

	client, err := NewClient(serverIp, serverPort)
	if err != nil {
		fmt.Println(">>>Error connecting to server...")
	}

	// open a goroutine to handle server's response message
	// handle server response
	go client.DealResponse()

	fmt.Println(">>>Success connecting to server...")

	// send to server
	client.Run()
}
