package main

import (
	"fmt"
	//	"sync"
)

/*
Golang并发编程之WaitGroup实现同步，什么是同步？就是两个goroutine之间相互等待
*/

//var wg sync.WaitGroup

func showMsg(i int) {
	//defer wg.Done() //goroutine结束就-1
	fmt.Println("Hello  Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		//wg.Add(1) // 循环一次，就启动一个goroutine就+1
		go showMsg(i)
	}
	//wg.Wait() //等待所有登记的goroutine都结束，为0才退出。
	//time.Sleep(time.Millisecond * 1)
	fmt.Println("end...")
}
