package main

import "fmt"

/*
一、Go语言没有面向对象的概念，但是可以使用结构体来实现，面向对象编程的一些特性，如：继承，组合等特性
二、语法：
type struct_var_type struct
member definition; //成员定义
member definition;
....
member definition;

三、未初始化的结构体成员都是零值，int=0,float=0.0 bool=false, string=nil
四、可以只对需要的成员进行初始化

*/
type person struct {
	id            int
	name          string
	age           int
	wechat, email string //同一类型可以合并到一行

}

func main() {

	var tom person
	fmt.Printf("tom: %v\n", tom)  //tom: {0  0  } ,为零值

	tom.id = 1
	tom.name = "tom"
	tom.age = 19
	tom.wechat = "110110"
	tom.email = "tom@n.com"
	fmt.Printf("tom: %v\n", tom)           //tom: {1 tom 19 110110 tom@n.com}
	fmt.Printf("tom.name: %v\n", tom.name) //tom.name: tom

	//匿名结构体,直接用变量定义个Struct直接使用
	var dog struct {
		id   int
		name string
	}
	dog.id = 1
	dog.name = "tiger"
	fmt.Printf("dog: %v\n", dog)  //dog: {1 tiger}

}
