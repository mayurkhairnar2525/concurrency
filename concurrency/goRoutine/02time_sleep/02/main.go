package main

import (
	"fmt"
	"time"
)

func main (){
	go	SayHello()
	fmt.Println("hello from the main")

	time.Sleep(100*time.Millisecond)
}

func SayHello (){
	fmt.Println("Hello from the function sayhello")
}
