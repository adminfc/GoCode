package main

import (
	"fmt"
	"sync"
	"time"
)

/*
除了可以使用channel实现同步完，还可以使用Mute互斥锁的方式来实现同步。
*/
var m int = 100
var lock sync.Mutex
var wt sync.WaitGroup

func add() {
	defer wt.Done()
	lock.Lock()
	m += 1
	time.Sleep(time.Second)
	lock.Unlock()
	fmt.Printf("m: %v\n", m)
}
func sub() {
	defer wt.Done()
	lock.Lock()
	m -= 1
	time.Sleep(time.Second)
	lock.Unlock()
	fmt.Printf("m: %v\n", m)
}
func main() {
	for i := 0; i < 5; i++ {
		go add()
		wt.Add(1)
		go sub()
		wt.Add(1)
	}
	wt.Wait()
	fmt.Printf("m end is: %v\n", m)
}
