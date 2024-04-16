package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	port := flag.String("port", "8080", "the port number")
	host := flag.String("host", "", "domain FQDN")
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello, server!"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received: %s", string(buf[:n]))
}
