package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.String("port", "8080", "the port number")
	host := flag.String("host", "", "domain FQDN")
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		time.Sleep(1000 * time.Millisecond)
		conn.Close()
	}()

	requests := []string{
		"ougmcim|SIGN_IN|janedoe\n",
		"iwhygsi|WHOAMI\n",
		"cadlsdo|SIGN_OUT\n",
		"asdasas|WHOAMI\n",
	}

	reader := bufio.NewReader(conn)

	for _, request := range requests {
		_, err = conn.Write([]byte(request))
		if err != nil {
			log.Fatal(err)
		}

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received: %s", response)
	}
}
