package main

import (
	"log"
	"net"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server //which server is current user belong to
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	go user.ListenMessage()

	return user
}

// call when user is online
func (u *User) Online() {
	// add user to online map
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()

	// broadcast current user online message to server
	u.server.Broadcast(u, "Online")
}

// call when user is offline
func (u *User) OffLine() {
	// remove user from the user list
	u.server.mapLock.Lock()
	if _, exist := u.server.OnlineMap[u.Name]; exist {
		delete(u.server.OnlineMap, u.Name)
		u.server.mapLock.Unlock()

		u.server.Broadcast(u, "Offline")
	} else {
		log.Printf("User not exist")
	}
}

// handle user message
func (u *User) DoMessage(msg string) {
	u.server.Broadcast(u, msg)
}

func (u *User) ListenMessage() {
	for {
		msg := <-u.C

		u.conn.Write([]byte(msg + "\n"))
	}
}
