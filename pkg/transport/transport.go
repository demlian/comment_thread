package transport

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

// Connection represents a TCP connection.
type Connection struct {
	Conn net.Conn
}

func (c *Connection) Write(responseID, responseData string) {
	n, err := c.Conn.Write([]byte(responseID + responseData + "\n"))
	if err != nil {
		log.Println(err)
	}
	log.Printf("wrote %d bytes", n)
}

// ComputeHash computes the hash of the TCP 5-tuple.
func (c *Connection) ComputeHash() (string, error) {
	localAddr := c.Conn.LocalAddr().String()
	remoteAddr := c.Conn.RemoteAddr().String()
	localPort, err := getPort(localAddr)
	if err != nil {
		return "", err
	}
	remotePort, err := getPort(remoteAddr)
	if err != nil {
		return "", err
	}

	ft := fmt.Sprintf("%s:%d-%s:%d", localAddr, localPort, remoteAddr, remotePort)
	hashBytes := sha256.Sum256([]byte(ft))
	hash := hex.EncodeToString(hashBytes[:])

	return hash[:], nil
}

// getPort extracts the port number from an address string.
func getPort(addr string) (int, error) {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		return 0, err
	}
	return net.LookupPort("tcp", port)
}
