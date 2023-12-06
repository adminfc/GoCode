package main

import "fmt"

/*
值接收者是一个拷贝，是一个副本。指针接收者，传递的是指针


*/

type Pet interface {
	eat()
}

type Dog struct {
	name string
}

/* func (dog Dog) eat()  {  //测试一
	fmt.Printf("dog: %p\n", &dog)
	fmt.Println("dog eat...")
	dog.name ="黑虎"
} */

func (dog *Dog) eat() { //测试二，将Pet接口改为指针接收者
	fmt.Printf("dog: %p\n", &dog)
	fmt.Println("dog eat...")
	dog.name = "黑虎"
}

func main() {
	/* 	fmt.Println("-----------测试一，复制了一份，dog name没有改变，外面的dog变量没有被修改----------")
	   	dog:=Dog{name : "花花"}
	   	fmt.Printf("dog: %p\n", &dog)
	   	dog.eat()
	   	fmt.Printf("dog: %v\n", dog)
	   	/*
	   	dog: 0xc000050250
	   	dog: 0xc000050260
	   	dog eat...
	   	dog: {花花}

	*/

	fmt.Println("-----------测试二，将Pet接口改为指针接收者----------")
	dog := &Dog{name: "花花1"}
	fmt.Printf("dog: %p\n", dog)
	dog.eat()
	fmt.Printf("dog: %v\n", dog)
	/*
		dog: 0xc000050250
		dog: 0xc00000a030
		dog eat...
		dog: &{黑虎}
	*/

}
