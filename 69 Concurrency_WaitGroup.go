package main

import (
	"fmt"
	"sync"
)

/*
Golang并发编程之WaitGroup实现同步，什么是同步？就是两个goroutine之间相互等待
如这个示例中，你可以用for循环顺序输出0-9,但是单线程的，如果是要并发，那就需要协程，那协程要执行完，就必须用到waitgroup
waitgroup.add(-1)  = waitgroup.done()
waitgroup.add (1),表示不为0则不会退出，需要等待。
waitgroup.wait 表示等待最后为0才退出。
*/

var wg sync.WaitGroup

func showMsg(i int) {
	defer wg.Done() //goroutine结束就-1
	fmt.Println("Hello  Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 循环一次，就启动一个goroutine就+1
		go showMsg(i)
	}
	wg.Wait() //等待所有登记的goroutine都结束，为0才退出。
}
