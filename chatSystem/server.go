package main

import (
	"fmt"
	"net"
)

//def localhost "127.0.0.1"
type Server struct {
	Ip   string
	Port int
}

//创建一个server
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

//
func (this *Server) Handler(conn net.Conn) {
	//TODO：当前链接的业务代码
	fmt.Println("链接建立成功")
}

//启动Server
func (this *Server) Start() {
	//创建socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	defer listener.Close()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err : ", err)
			continue
		}
		// do handler
		go this.Handler(conn)
	}
	// close listen socket
}
