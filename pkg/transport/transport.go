package transport

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
)

// Connection represents a TCP connection.
type Connection struct {
	Conn   net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func (c *Connection) Init() {
	c.Reader = bufio.NewReader(c.Conn)
	c.Writer = bufio.NewWriter(c.Conn)
}

func (c *Connection) ReadMessage() (string, error) {
	data, err := c.Reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			log.Println("client closed the connection")
		}
		return "", err
	}
	return data, nil
}
func (c *Connection) Write(responseID, responseData string) {
	data := []byte(fmt.Sprintf("%s%s\n", responseID, responseData))
	_, err := c.Writer.Write(data)
	if err != nil {
		log.Println(err)
		return
	}
	err = c.Writer.Flush()
	if err != nil {
		log.Println(err)
	}
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
