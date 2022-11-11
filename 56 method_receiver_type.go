package main

import "fmt"

/*
结构体实例中，有值类型和指针类型，那么方法的接收者是结构体，那么也有值类型和指针类型、
区别就是接收者是否复制结构体副本，值类型就是复制，指针类型不复制。

*/
type Person struct {
	name string
}

func showPerson(per Person) {
	fmt.Printf("per: %p\n", &per)
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func showPerson2(per *Person) {
	fmt.Printf("per: %p\n", per)
	per.name = "Kite"
	fmt.Printf("per: %v\n", per)
}

func main() {
	p1 := Person{name: "Tom"}
	p2 := &Person{name: "Tom"}

	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p2: %T\n", p2)
	fmt.Println("-----------------------")
	/*
		p1: main.Person
		p2: *main.Person
	*/

	p3 := Person{name: "Tom"}
	fmt.Printf("p3: %p\n", &p3)
	showPerson(p3)
	fmt.Printf("p3: %v\n", p3)

	fmt.Println("-----------------------")

	p4 := &Person{name: "Tom"}
	fmt.Printf("p4: %p\n", p4)
	showPerson2(p4)
	fmt.Printf("p4: %v\n", p4)
	/*
	   p4: 0xc0000502b0
	   per: 0xc0000502b0
	   per: &{Kite}
	   p4: &{Kite}

	*/
}
