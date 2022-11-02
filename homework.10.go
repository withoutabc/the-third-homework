package main

import (
	"fmt"
	"sync"
)

var i int64
var x int64
var wg sync.WaitGroup

func add() {
	ch := make(chan int64, 1)
	for i = 0; i < 50000; i++ {
		ch <- i + 1 //把i放入管道
		x = <-ch
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
