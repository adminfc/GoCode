package main

import "fmt"

//key:value,使用make或者map来定义map

//创建map
func test1() {
	var s1 map[string]string
	s1 = make(map[string]string)
	fmt.Printf("s1: %T\n", s1)
}

func test2() {
	s2 := map[string]string{"Name": "Brade", "age": "18"}
	fmt.Println(s2)

	s3 := make(map[string]string)
	s3["name"] = "Brade"
	s3["age"] = "13"
	s3["email"] = "brade.y.zhu@Newegg.com"

	fmt.Printf("s3: %v\n", s3)

	delete(s3, "age")
	fmt.Printf("s3: %v\n", s3)
}

//访问map
func test3() {

	s4 := make(map[string]string)
	s4["name"] = "Brade"
	s4["age"] = "13"

	fmt.Println(s4["name"], s4["age"])
}

//判断是否存在
func test4() {
	s5 := map[string]string{"Name": "Brade", "age": "19"}
	value1, ok1 := s5["Name"]
	value2, ok2 := s5["xx"]
	fmt.Println(value1, ok1)
	fmt.Println(value2, ok2)

}

//遍历map
func test5() {
	s6 := map[string]string{"Name": "Brade", "age": "17"}
	for key, value := range s6 {
		fmt.Println(key, value)
	}

}

func main() {
	//test1()
	test2()
	//test3()
	//test4()
	//test5()
}
