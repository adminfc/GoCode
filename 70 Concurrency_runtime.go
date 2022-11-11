package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
runtime包里面定义了一些协程相关的api：
1、runtime.Gosched()  让出CPU时间片，重新等待安排任务。
2、runtime.goexit()   退出当前协程
*/

// 示例1，runtime.goSched.
func show(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}

// 示例2  runtime.Goexit()
func showexit(msg string) {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit()
		}
		fmt.Printf("i: %v\n", i)
	}
}

func main() {
	//示例1代码
	/* 	go show("java")          //示例1：子协程来执行
	   	//主协程里面加了runtime.Gosched，它会主动把CPU时间片让出来，让子协程先运行。
	   	for i := 0; i < 2; i++ {
	   		runtime.Gosched() //切一下，再次分配任务，可注释一下看不同的结果
	   		fmt.Println("golang")
	   	}
	*/

	/*
		java
		java
		golang
		golang

	*/

	//示例2代码
	go showexit("Hi Google") //示例2子协程，大于5就退出协程了
	time.Sleep(time.Second)
	/*
		i: 0
		i: 1
		i: 2
		i: 3
		i: 4
	*/
}
