/**
 * Follow the code sample at Systembash
 * @see https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
 *
 * Run the server first with `go run tcp-server.go`. Then run the client with `go run tcp-client.go`
 */

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghi jklmnopqrs tuvwxyz ABCDEFGHI JKLMNOPQRST UVWXYZ 0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Println("Lauching server...")

	// Listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// Accept connection on port
	conn, _ := ln.Accept()

	// Run loop forever (or until Ctrl + C)
	for {
		// Will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// Output message received
		fmt.Print("Message received:", string(message))

		// Sample process for string received
		newmessage := randSeq(30)
		fmt.Println("Reply a random message:", newmessage, "\n")

		// Send new String back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
