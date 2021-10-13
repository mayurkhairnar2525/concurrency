package main

import (
	"fmt"
	"time"
)

func main (){
go	SayHello()
	time.Sleep(100*time.Millisecond)
}

func SayHello (){
	fmt.Println("Hello from the function sayhello")
}