package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var Counter = 0

func main (){
	for i:=0;i<10;i++{
		wg.Add(2)

		go sayhello()
		go increament()
	}
	wg.Wait()
}

func sayhello (){
	fmt.Printf("Hello #%v\n",Counter)
	wg.Done()
}

func increament (){
	Counter ++
	wg.Done()
}