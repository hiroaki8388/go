package main

import (
	"io"
	"log"
	"net"
)

//   同時接続1
func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
		log.Print("Connect!!")
	}
}

//
func handleConn(c net.Conn) {
	io.Copy(c, c)
	aaa
	c.Close()
}
