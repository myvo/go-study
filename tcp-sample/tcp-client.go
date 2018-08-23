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
	"net"
	"os"
)

func main() {
	// Connect to this Socket
	conn, _ := net.Dial("tcp", ":8081")
	for {
		// Read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// Send to Socket
		fmt.Fprintf(conn, text+"\n")

		// Listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message from server:", message)
	}
}
