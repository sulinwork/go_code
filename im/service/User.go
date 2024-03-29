package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	user := &User{
		Name: conn.RemoteAddr().String(),
		Addr: conn.RemoteAddr().String(),
		C:    make(chan string),
		conn: conn,
	}
	user.Listener()
	return user
}

func (this *User) Listener() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
