package main

import "fmt"

/*
一、结构体指针和普通的变量指针相同
二、可以使用new关键字来创建结构体指针
三、注意一下不同的方法，比如变量定义，结构体定义，new定义时成员的赋值方法
*/

func p_var() {
	name := "tom"
	var p_name *string
	p_name = &name
	fmt.Println(name)
	fmt.Println(p_name)
	fmt.Println(*p_name)
	/*
		通指针和变量输出示例：
		tom
		0xc000052250
		tom
	*/

}

func s_struct() {
	type person struct {
		id   int
		name string
	}
	tom := person{
		id:   101,
		name: "Tom",
	}
	var p_person *person
	p_person = &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)
	/*
		tom: {101 Tom}
		p_person: 0xc000008078
		p_person: {101 Tom}
	*/
}

func test() {
	type person struct {
		id   int
		name string
	}

	var tom = new(person) //使用new关键字来创建结构体指针
	tom.id = 102
	tom.name = "Jerry"
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("tom: %v\n", tom.name)
}

func main() {
	p_var()    //通指针和变量示例
	s_struct() //结构体指针示例
	test()

}
