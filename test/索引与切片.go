package main

import (
	"fmt"
)

func main() {
	var name string = "朱云飞"

	fmt.Println(name[0])   //230  //1. 索引获取字节
	fmt.Println(name[0:3]) //2. 切片获取字节区间

	//3. 手动循环获取所有字节
	for i := 0; i < len(name); i++ {
		fmt.Println(i, name[i])
	}

	//4. for range循环获取所有字符
	for index, item := range name {
		fmt.Println(index, item)
	}

	//5.转换成rune集合
	daalist := []rune(name)
	fmt.Println(daalist[0], string(daalist[0]))
}
