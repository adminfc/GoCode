package main

import "fmt"

/*
一、go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无需拷贝数据。
类型指针不能进行偏移和运算。
go语言中的指针操作非常简单，只城赅住两个符号： &（取地址）和*（根据地址取值）
go语言中的值类型都有对应的指针类型，比如*int, *int64, *string,*bool

格式：
var var_name *var-type

二、指向数据的指针
表示数组中的每个元素的类型是指针类型
格式：
var ptr [MAX]*int

*/

func arr() {
	a := []int{1, 3, 5}
	var ptr [3]*int
	fmt.Println(ptr) //[<nil> <nil> <nil>]
	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

}

func main() {
	/* 	var i *int
	   	fmt.Println(i) //<nil>

	   	var s int = 100
	   	i = &s
	   	fmt.Println(*i) //100
	   	fmt.Println(i)  //0xc0000180c0
	   	fmt.Println(&s) //0xc0000180c0
	   	fmt.Println(&i) //0xc00000a028 */
	arr()
}
