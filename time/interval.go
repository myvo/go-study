package main

import "time"
import "fmt"

func main() {
	// Interval in every 1 second.
	tickChan := time.NewTicker(time.Second * 1).C

	<-tickChan
	for t := range tickChan {
		fmt.Println("Tick at", t)
	}
}
