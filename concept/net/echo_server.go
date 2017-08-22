package main

import (
	"log"
	"net"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Printf("failed to read from %v: %v", conn.RemoteAddr(), err)
			return
		}

		log.Printf("got from %v: %s\n", conn.RemoteAddr(), string(buf[0:n]))

		if _, err := conn.Write(buf[0:n]); err != nil {
			log.Printf("failed to write to %v: %v", conn.RemoteAddr(), err)
			return
		}
	}
}
