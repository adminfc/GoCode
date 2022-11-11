package main
import "fmt"
//高级函数及匿名函数

//示例1，函数作为参数使用
func sayHello(name string)  {
	fmt.Println("hello",name)
}

func test(name string,f func(string))  {
	f(name)
}


//示例2，匿名函数  func(参数列表)(返回值)，可以没有参数，没有返回值 
func test2()  {
	max := func (a int,b int) int{
		if a> b {
			return a
		}else{
			return b
		}
	}
i := max(1,11)
fmt.Println(i)
}



func main() {
	test("Brade",sayHello) //示例1输出结果hello Brade
	test2()  //输出最大值 
}