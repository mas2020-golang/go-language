package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Hey, I'm listening on port 8000...")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

// Manage the single connection in a separate go routine.
func handleConn(c net.Conn) {
	defer c.Close()
	fmt.Println("new client connection: ESTABLISHED")
	for {
		// write current time to the client
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if ok := testError(err); ok {
			break
		}
		time.Sleep(1 * time.Second)

		// read from connection until it reaches '\n' single char
		br := bufio.NewReader(c)
		dataBytes, err := br.ReadBytes('\n')
		if ok := testError(err); ok {
			break
		}
		// write data sent by the client back to the client
		_, err = io.WriteString(c, string(dataBytes))
		if ok := testError(err); !ok {
			break
		}
	}
	fmt.Println("client connection: DISCONNECTED")
}

func testError(err error) bool {
	if err != nil {
		return true
	}
	return false
}
