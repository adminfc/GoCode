package main

import (
	"fmt"
)

func main() {

	//方式一，先声明再赋值，初始化为0
	var num [3]int
	num[0] = 111
	num[2] = 222
	fmt.Println(num)

	//方式二 声明+赋值
	var names = [2]string{"M", "A"}
	fmt.Println(names)

	//方式三 声明+赋值+指定位置
	var ages = [3]int{0: 1, 1: 99, 2: 99}
	fmt.Println(ages)

	//方式四 不指定个数

	var AAs = []string{"33"}
	fmt.Println(AAs)

	//声明数组并初始化，返回的是指针
	BBs := new([3]int)
	fmt.Println(BBs)

	/*
		1。数组的元素是可以更改的，但类型和长度不可修改。name[1] = "test"
		2.数组可以拷贝，
		name1 := [2]string{"A","B"}
		name2 := name1
		name1[1] = "C"
		fmt.println(name1,name2)  //["A","C"]  ["A","B"]
	*/

	//数组的长度索引和切片

	//长度
	mooing := []string{"朱A", "云", "飞"}
	fmt.Println(len(mooing))
	//索引
	fmt.Println(mooing[2])
	//切片
	fmt.Println(mooing[0:2])
	//循环
	for i := 0; i < len(mooing); i++ {
		fmt.Println(i, mooing[i])
	}

	for num, item := range mooing {
		fmt.Println(num, item)
	}

	//数组嵌套
	var data [2][3]int
	data[0] = [3]int{1, 2, 3}
	fmt.Println(data[0][2])
}
