package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//List of online users
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	//msg broadcast channel
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// listen message chan in server, once received, send to all users
func (s *Server) ListenMessager() {
	for {
		msg := <-s.Message
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

func (s *Server) Broadcast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)

	s.Message <- sendMsg
}

// handle user online
func (s *Server) Handler(conn net.Conn) {
	// create new user
	user := NewUser(conn, s)

	user.Online()

	// receive message from the client
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				log.Println("conn read err:", err)
				return
			}
			if n == 0 {
				user.OffLine()
				return
			}

			// extract user's message and remove `\n` at the end
			msg := string(buf[:n-1])

			// broadcast the msg
			user.DoMessage(msg)
		}
	}()
}

func (s *Server) Start() {
	// socket listen
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}

	defer listen.Close()

	//start messager listen process
	go s.ListenMessager()

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
