package main

//闭包并不复杂，只要牢记  ： 闭包=函数+引用环境。

import "fmt"

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}

}

func main() {
	f := add() //变量f是一个函数，并且它引用了其外部作用域中的x变量，此时f就是一个闭包，在f的生命周期内，变量x也一直有效。
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)
	r = f(30)
	fmt.Printf("r: %v\n", r)
	fmt.Println("---------------------") //在没有重新定义之前，值会被保留，重新定义后，值将不在保留

	f1 := add()
	r = f1(100)
	fmt.Printf("r: %v\n", r)

}
