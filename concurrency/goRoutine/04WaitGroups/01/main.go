package main

import "fmt"

func runner1(){
	fmt.Println("runner 1")
}

func runner2(){
	fmt.Println("runner 2")
}

func execute(){
	runner1()
	runner2()
}
func main (){
	execute()
}
