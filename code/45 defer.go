package main

import "fmt"

/*
特性
1.关键字defer用于注册延迟调用。
2. 这些调用直到return前才被执行，因此，可以用来做资源清理
3. 多个defer语句，先进后出的方式执行
4. defer语句中的变量，在defer声明时就决定了

用途：
1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放

*/
func main() {
	fmt.Println("Start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("end")
}

/*
Start
end
step3
step2
step1
*/
