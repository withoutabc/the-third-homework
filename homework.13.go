package main

import (
	"fmt"
)

var i int

func main() {
	over := make(chan bool)
	go func() {
		for i = 0; i < 10; i++ {
			fmt.Println(i)
			if i == 9 {
				over <- true
			}
		}
	}()
	<-over
	fmt.Println("over!!!")

}

//for循环,if语句应写在协程里
//不知道这样改是否符合要求
