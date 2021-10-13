package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg=sync.WaitGroup{}
var Counter = 0
var m = sync.RWMutex{}

func main (){
runtime.GOMAXPROCS(100)
for i:=0;i<10;i++{
	wg.Add(2)
	go sayhello()
	go Increament ()
}
wg.Wait()
}
func sayhello(){
	m.RLock()
	fmt.Println("Hello World",Counter)
	m.RUnlock()
	wg.Done()
}

func Increament(){
	m.Lock()
	Counter ++
	m.Unlock()
	wg.Done()
}
