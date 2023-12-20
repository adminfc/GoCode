package main

import "fmt"

/*
定义：
接口像一个公司的领导，他会定义一些通用规范，只设计规范，而不实现规范
go语言的接口，是一定新的类型定义，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
语法和方法非常类似：
//定义接口
type interface_name interface {
	method_name1 [return_type]
	method_name2 [return_type]
	.....
	method_namen [return_type]
}
//定义结构体：
type struct_name struct{
	//定义一些变量
}

//实现接口方法：
func (struct_name_var struct_name) method_name1() [return_type]{
	//方法实现
}
func (struct_name_var struct_name) method_namen() [return_type]{
	//方法实现
}
*/

//定义一个USB接口，有write和read两个方法，再定义一个computer和一个Mobile来实现这个接口

type USB interface {
	read()
	write()
}

type Computer struct {
}
type Mobile struct {
}

func (c Computer) read() {
	fmt.Println("Computer read...")
}

func (c Computer) write() {
	fmt.Println("Computer write...")
}

func (c Mobile) read() {
	fmt.Println("Mobile read...")
}

func (c Mobile) write() {
	fmt.Println("Mobile write...")
}

func main() {
	c := Computer{}
	c.read()
	c.write()
	fmt.Println("-------------------------")
	m := Mobile{}
	m.read()
	m.write()

}
