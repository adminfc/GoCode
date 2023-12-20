package main

import "fmt"

/*
递归：函数内部自己调用自己的函数叫递归函数
1. 必须先定义函数退出条件，没有退出条件，递归就是死循环。
2. 可能会产生一大堆goroutine,也有可能出现栈空间内存溢出问题。
*/

func f1() {
	//for循环求5的阶乘，5X4X3X2X1
	var s int = 1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

func f2(a int) int {
	//用递归函数来算阶乘
	if a == 1 {
		//1.退出条件
		return 1
	} else {
		return a * f2(a-1) //自我调用
	}
}

func f(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return f(n-1) + f(n-2)
}

func main() {
	f1()
	fmt.Println(f2(5))
	r := f(5)
	fmt.Printf("r: %v\n", r)
}
