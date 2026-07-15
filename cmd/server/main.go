package main

import "flag"

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "setting server IP address")
	flag.IntVar(&serverPort, "port", 8080, "setting server port number")
}

func main() {
	flag.Parse()

	server := NewServer(serverIp, serverPort)
	server.Start()
}
