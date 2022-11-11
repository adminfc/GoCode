package main

import "fmt"

/*
Go语言类型语法：
type newType Type
type myInt int

GO语言类型别名
type NewType = Type
type MyInt2 = int

区别：
1。类型定义相当于定义了一个全新的类型，与之前的类型不同；但是类型别名并没有定义一个新的类型，而是使用一个别名来替换之前的类型
2。类型别名只会在代码中存在，编译完之后并不会存在
3。因为类型别名和原来的类型是一样的，所以原来类型所惹的人方法，别名中也可以调用，但是如果是重新定义一个类型，那么不可以调用之前的方法

*/
func main() {

	type MyInt int
	var i MyInt
	i = 100
	fmt.Printf("i: %T,%v\n", i, i)

	type MyInt2 = int
	var ii MyInt2
	ii = 100
	fmt.Printf("i: %T,%v\n", ii, ii)
}

/*
i: main.MyInt,100
i: int,100
*/
