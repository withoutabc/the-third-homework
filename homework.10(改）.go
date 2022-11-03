package main

import (
	"fmt"
	"sync"
)

var i int64
var x int64
var wg sync.WaitGroup
var ch = make(chan int64, 1)

func add() {
	ch <- 555
	for i = 0; i < 50000; i++ {
		x = x + 1
	}
	<-ch
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
