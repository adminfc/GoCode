package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
 runtime.goMaxProcs  //设置最大的CPU核心数，默认情况下最多的核心数
*/

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B", i)
	}
}
func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU()) //获取当前主机的核心数
	runtime.GOMAXPROCS(2)                                  //修改为1时，要么A先要么B先。如果是设置为2就是交替
	go a()
	go b()
	time.Sleep(time.Second)
}
