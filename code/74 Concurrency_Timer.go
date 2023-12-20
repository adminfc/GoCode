package main

import (
	"fmt"
	"time"
)

/*
Timer:定时器，可以实现一些定时操作，其内部也是通过Channel来实现的。
*/
func Monitor() { //每隔一秒输出一次
	ticker := time.NewTicker(time.Second)
	for _ = range ticker.C {
		fmt.Println("ticker....")

	}
}

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Printf("t1: %v\n", t1)
	t2 := <-timer1.C //阻塞的，指定时间到了

	fmt.Printf("t1: %v\n", t2)

	Monitor()

	time.Sleep(time.Second * 3) //延迟3s
	time.After(time.Second * 3)
	time.AfterFunc(time.Second*3, func() { fmt.Printf("time.Now(): %v\n", time.Now()) })
	timer1.Reset(time.Second * 1)
}
