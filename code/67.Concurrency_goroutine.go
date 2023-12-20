package main

import (
	"fmt"
	"time"
)

/*
并发编程之协程：
goroutines是并发运行的函数，创建一个协程非常简单，就是在一个任务函数前面添加一个go关键字
go task()
*/

func show(msg string) {
	for i := 1; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)

	}
}

func main() {
	go show("Java")        //启动了一条协程
	show("golang")         //如果这里也加上go，多一个协程，就不会显示了。或者加一个sleep等待
	fmt.Println("end....") //主程序结束，协程也就结束了
}
