package main

import "fmt"

//普通函数，匿名函数，方法
//经常使用返回值判断是否执行成功，是否有错误信息，如return value、exists,return value、ok,return value、err
/*
func function_name([parameter list])[return_type]  {
	函数体
}
*/

//示例求和
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

//无参数有返回值函数。可以没有参数，但是参数一定要有数据类型
func test1() (name string, age int) {
	name = "Brade"
	age = 18
	return
}

//如果是map,slice,interface,channel这些数据本身就是指针类型，传值的时拷贝的也是指针，
//拷贝后的参数仍然指向底层数据结构，修改他们可能会影响外部数据结构的值
func test2(s []int)  {
	s[0] =  1000
}



func main() {
/* 	r := sum(2, 3)
	fmt.Printf("r: %v\n", r) */

/* 	name, age := test1()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age) */

	s := []int{1,2,3}
	test2(s)
	fmt.Printf("s: %v\n", s)  //发现s[0]仍然是1000

}
