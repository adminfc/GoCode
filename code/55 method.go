package main

import "fmt"

/*
go中的方法是一种特殊的函数，定义于struct之上（与struct关联，绑定），被称为struct的授受者(receiver)
通俗的讲，方法就是有接收者的函数

语法格式：
type myType struct{}
func (recv myType) my_method(para) return_type{}   //在func和函数名之间加了个(recv myType)，后面就和函数语法一样了
func (recv *myType) my_method(para) return_type{}

myType: 结构体名称
my_method: 方法名
para: 参数列表
return_type: 返回值类型

从语法格式可以看出，一个方法和一个函数非常相似，多了一个授受类型。
注意：
1。方法receiver type不一定是struct类型，可以是slice,map,channel,func
2。struct结合它的方法就等价于面向对象中的类

*/

type Person struct {
	name string
}

func (per1 Person) eat() { //(per Person) per1就是receiver
	fmt.Println(per1.name + " eating....")
}

func (per1 Person) sleep() {
	fmt.Println(per1.name + " sleeping....")
}

func main() {
	var per Person
	per.name = "Tom"
	per.eat()
	per.sleep()

	/*
			Tom eating....
		    Tom sleeping....
	*/
}
