package main

import "fmt"

//函数类型与函数变量
//type fun func(int,int) int ,定义了一个fun函数类型，这个函数接收两个int参数，返回一个int

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	type f1 func(int, int) int  //定义一个f1的函数类型，把sum和max值给它
	var ff f1   //定义一个ff的变量，类型为f1
	ff = sum
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)

	ff = max
	r = ff(1, 5)  //上面已经定义过变量了，所以这里直接=
	fmt.Printf("r: %v\n", r)

}
