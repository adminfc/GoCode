package main

import (
	"fmt"
	"reflect"
)

func main() {
	//定义结构体实例参数
	s := student{
		Name: "Lilei",
		Age:  18,
	}
	testReflect(&s)
	fmt.Printf("s: %v\n", s)
}

//利用一个函数，函数的参数定义为空接口（空接口没有方法，可以理解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋给空接口）

func testReflect(i interface{}) {
	//将i转换成reflect.value类型
	val := reflect.ValueOf(i)
	fmt.Printf("val: %v\n", val)

	n := val.Elem().NumField()
	fmt.Printf("n: %v\n", n)

	//修改字段的值
	val.Elem().Field(0).SetString("朱朱")

}

// 定义一个结构体并绑定方法
type student struct {
	Name string
	Age  int
}
