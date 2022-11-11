package main

import "fmt"

/*
go结构体可以像普通变量一样，作为函数的参数，传递给函数
1。直接传递结构体，这是一个副本拷贝，在函数内部不会改变外面结构体内容。
2。传递结构体指针，这时在函数内部，能改变外部结构体内容。
*/

type Person struct {
	id   int
	name string
}


func showPerson(person Person) {
	person.id = 1
	person.name = "kite"
	fmt.Printf("person: %v\n", person)
}

func main() {

	person := Person{1, "tom"}
	fmt.Printf("person: %v\n", person)  //函数外面没有改变结构体内容。
	fmt.Println("----------------")
	showPerson(person)                  //函数内部改变了结构体内容
	fmt.Println("----------------")
	fmt.Printf("person: %v\n", person)  //函数外面没有改变结构体内容。

	/*
	person: {1 tom}
	----------------
	person: {1 kite}
	----------------
	person: {1 tom}

	可以看出，函数内部改变了结构体内容，函数外面并没有被改变
	*/

}
