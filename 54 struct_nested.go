package main

import "fmt"

/*
go没有面向对象的编程思想，也没有继承关系，但是可以通过结构体嵌套来实现这种效果。
示例：有一个Person的结构体，这个人还养了一条dog结构体
*/
type Dog struct {
	name  string
	color string
	age   int
}

type P_person struct {
	dog  Dog
	name string
	age  int
}

func main() {
	var tom P_person
	tom.dog.name = "小花"
	tom.dog.color = "黑"
	tom.dog.age = 1
	tom.name = "Tom"
	tom.age = 20
	fmt.Println(tom) //{{小花 黑 1} Tom 20}

}
