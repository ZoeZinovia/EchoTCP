package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// open a connection to local host on port 7
	conn, err := net.Dial("tcp", ":7")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(conn, os.Stdin) // copy input from console to conn variable
	if err != nil {
		log.Fatal(err)
	}
	tcpconn, ok := conn.(*net.TCPConn) // tell the server that client is finished writing
	if ok {
		tcpconn.CloseWrite()
	}
	_, err = io.Copy(os.Stdout, conn) // copy received response to console output
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Close() // close the connection
	if err != nil {
		log.Fatal(err)
	}
}
