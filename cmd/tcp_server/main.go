package main

import (
	"bufio"
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
	reader := bufio.NewReader(conn.Conn)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("Client closed the connection")
				break
			}
			log.Println(err)
			continue
		}
		connHash, err := conn.ComputeHash()
		if err != nil {
			log.Println(err)
		}

		response, err := protocol.HandleRequest(connHash, data)
		if err != nil {
			log.Println(err)
			continue
		}
		conn.Write(response.ID, response.Data)
	}
}
