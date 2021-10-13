package main

import (
	"fmt"
	"sync"
)

func runner1(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("runner 1")
}

func runner2(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("runner 2")
}

func execute(){
	wg :=new(sync.WaitGroup)
	wg.Add(2)
	go 	runner1(wg)
	go 	runner2(wg)
	wg.Wait()
}
func main (){
	execute()

}


