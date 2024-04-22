package main

import (
	"flag"
	"io"
	"log"
	"net"

	"github.com/demlian/comment_thread/pkg/protocol"
	"github.com/demlian/comment_thread/pkg/transport"
)

func main() {
	port := flag.String("port", "8080", "the port number")
	flag.Parse()
	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Server listening on %s:%s", listener.Addr().(*net.TCPAddr).IP, *port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("connection failed:", err)
			continue
		}
		log.Println("Accepted connection from:", conn.RemoteAddr())

		connection := transport.Connection{Conn: conn}
		go handleConnection(connection)
	}
}

func handleConnection(conn transport.Connection) {
	defer conn.Conn.Close()
	conn.Init()

	for {
		data, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				log.Println("client closed the connection")
				break
			}
			log.Println(err)
			continue
		}

		// Compute the hash of the connection 5-tuple to identify the user.
		uuid, err := conn.ComputeHash()
		if err != nil {
			log.Println(err)
		}

		responseString, err := protocol.HandleRequest(uuid, data)
		if err != nil {
			log.Println(err)
			continue
		}
		conn.Write("", responseString)
	}
}
