package main

import (
	"fmt"
	"time"
)


func Portal1 (channel1 chan string){
	time.Sleep(3*time.Second)
	channel1 <- "welcome to channel 1"
}
func Portal2 (channel2 chan string){
	time.Sleep(4*time.Second)
	channel2 <- "welcome to channel 1"
}

func main (){
	r1 := make (chan string)
	r2 := make(chan string)

	go Portal1(r1)
	go Portal2(r2)

	select {
	case op1 := <-r1:
		fmt.Println(op1)
	case op2 := <-r2:
		fmt.Println(op2)
		//default:
		//	fmt.Println("Not found")
		//}
	}
}