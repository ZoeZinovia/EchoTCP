// Created by Zoe Zinovia

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":7") // start listener on port 7
	if err != nil {
		log.Fatal(err)
	}

	// Infinite loop to listen for incoming requests
	for {
		conn, err := ln.Accept() // Accept an incoming request
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Echo request received")
		go echo(conn) // concurrently process each connection
	}
}

func echo(conn net.Conn) {
	_, err := io.Copy(conn, conn) // copy incoming connection content to outgoing connection content
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
