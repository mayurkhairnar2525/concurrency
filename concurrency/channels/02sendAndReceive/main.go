package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Println("Sending value to channel")
	go send(ch)
	fmt.Println("Receiving from channel")
	go receive(ch)
	time.Sleep(time.Second * 1)
}
func send(ch chan int) {
	ch <- 1
}
func receive(ch chan int) {
	val := <-ch
	fmt.Printf("Value Received=%d in receive function\n", val)
}
