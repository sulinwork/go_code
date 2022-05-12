package main

import (
	"fmt"
	"net"
)

type Server struct {
	ip   string
	port int
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip:   ip,
		port: port,
	}
}

func (this *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.ip, this.port))
	if err != nil {
		fmt.Print("异常==")
		return
	}
	defer listen.Close()
	for {
		if accept, err := listen.Accept(); err == nil {
			go this.messageHandler(accept)
		}
	}
}

func (this *Server) messageHandler(accept net.Conn) {
	NewUser(accept)
}
