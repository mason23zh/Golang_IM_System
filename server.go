package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NweServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (s *Server) Handler(conn net.Conn) {
	fmt.Println("Connection Established")
}

func (s *Server) Start() {
	// socket listen
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}

	defer listen.Close()

	for {
		// accept
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// handler
		s.Handler(conn)
	}
}
