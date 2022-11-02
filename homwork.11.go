package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int, 1)
var wg sync.WaitGroup

func count1() {
	for i := 0; i < 100; i++ {
		i++
		fmt.Printf("%d ", i)
		ch <- 555 //先在这里存放一个数据，等待25行取出来
		<-ch      //取出27存入的数据
	}
	wg.Done()
}
func count2() {
	for i := 1; i < 100; i++ {
		i++
		<-ch //取出16行存入管道的值
		fmt.Printf("%d ", i)
		ch <- 666 //存入一个数据，等待17行取出
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go count1()
	go count2()
	wg.Wait()

}

//之前没用管道试了挺多次，它会随机选择count1和count2，就打印不出来顺序。
//查询到了管道的同步功能。
//主要是让count1和count2里管道同步 16和25同步  17和27同步
//大概是这样，理解的还不是很好
