/**
 * Convert hex string to []byte
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		// Read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Hex string: ")
		flagStrHex, _ := reader.ReadString('\n')

		var rs string
		hexArr := strings.Split(flagStrHex, ":")
		for _, v := range hexArr {
			rs += "0x" + v + ", "
		}

		fmt.Println("Result:", rs)
	}
}
