package main

import (
	"flag"
	"io"
	"log"
	"net"
)

func main() {
	port := flag.String("port", "8080", "the port number")
	flag.Parse()
	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	n, err := io.Copy(conn, conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Received and sent %d bytes from/to %s", n, conn.RemoteAddr().String())
}
