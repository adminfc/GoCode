package main

import "fmt"

/*
用途：可以实现包级别的一些初始化操作
特性：
1. init先于main函数执行，不能被调用
2. init没有输入参数，返回值
3. 每个包可以有多个init函数
4. 包的每个源文件也可以用多个Init函数，这点比较特殊
5. 同一个包的init执行顺序，没有明确定义，编程时不要依赖这个顺序
6. 不同包的init函数按照包的导入的依赖关系决定执行顺序

初始化顺序 ：变量初始化>init()>main()
*/
func init() {
	fmt.Println("init function")
}

var i int = initvar()

func initvar() int {
	fmt.Println("initvar...")
	return 100
}

func main() {
	fmt.Println("main function")

}

/*
initvar...
init function
main function
*/
