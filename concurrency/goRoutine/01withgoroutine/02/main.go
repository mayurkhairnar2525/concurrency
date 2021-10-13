package main

import (
	"fmt"
)

func main (){
	go sayHello()
	fmt.Println("Hello from main function")

}

func sayHello(){
	fmt.Println("Hello from the function sayhello")
}
