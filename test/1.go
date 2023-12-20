package main

import (
	"fmt"
)

func main() {
	var pwd int

	for i := 3; i >= 1; i-- {
		fmt.Print("please input your password:")
		fmt.Scanln(&pwd)
		if pwd > 66 {
			fmt.Println("输入错误！")
			fmt.Println("你还有", i-1, "次输入机会")
		} else if pwd < 66 {
			fmt.Println("输入错误！")
			fmt.Println("你还有", i-1, "次输入机会")
		} else {
			fmt.Println("恭喜你")
			break
		}
	}

	/*
		 	var t2 string = "15"
			result, err := strconv.Atoi(t2)
			fmt.Println(result, err, reflect.TypeOf(result))
	*/
}
