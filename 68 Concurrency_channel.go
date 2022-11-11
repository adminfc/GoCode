package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
一、概念
GO的通道机制，用于在Goroutine（协程）之间共享数据，
当执行goroutine并发活动时，要在goroutine之间共享资源或数据，通道channel提供一种机制来保证同步交换
需要在声明channel时指定数据类型。数据在通道上传递：在任何给定时间只能有一个goroutine可以访问数据项，因此按照设计不会发生数据竞争。

二、类型
Channel分两种类型，无缓冲通道和缓冲通道。
无缓冲通道用于执行goroutine之间的同步通信，而缓冲通道用于执行异步通信。
无缓冲通道保证在发送和接收发生的瞬间执行两个goroutine之间的交换。缓冲通道不保证
Channel由make函数创建，该函数指定chan关键字和通道的元素类型。

三、语法：
unbuffered := make(chan int)  //整型无缓冲通道
buffered := make(chan int,10) // 整形有缓冲通道

使用内置函数make创建无缓冲和缓冲通道，make的第一个参数需要关键定chan， 然后是通道允许交换的数据类型

四、示例
将值发送到通道的代码块需要使用<-运算符：
语法：
goroutine1 := make(chan string, 5) //字符串缓冲通道
goroutine1 <- "China"  //通过通道发送字符串
以上是一个包含5个值的缓冲区的字符串类型的goroutine1通道，然后我们通过通道发送字符串China

从通道以旧换新值的代码块：
语法：
date := <-goroutine1 //从通道以旧换新字符串
<- 运行符附加到通道变量goroutine1的左侧，以接收来自通道的值

五、通道的发送和接收的特性
1。对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。
2。发送和接收操作中对元素值的处理都是不可分割的。
3。发送操作在完全完成之间会被阻塞，接收操作也是如此。

*/

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(100)
	fmt.Printf("value: %v\n", value)
	time.Sleep(time.Second * 3)
	values <- value
}

func main() {
	//从通道接收数据
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end...")

	/*
		value: 42
		wait...
		receive: 42
		end...
	*/

}
