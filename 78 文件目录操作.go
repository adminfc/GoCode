package main

import (
	"fmt"
	"os"
)

/*
D:\Program Files\Go\src\os
https://pkg.go.dev/std


*/
//创建文件

func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func createDir() {
	//err := os.Mkdir("test1", os.ModePerm)
	err := os.MkdirAll("test1/a/b", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除目录
func removeDir() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	err2 := os.RemoveAll("test1")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}

}

// 获得工作目录
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

//获得工作目录

func chWd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	os.Chdir("d:/")
	dir1, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir1)

	s := os.TempDir() //获得临时目录
	fmt.Printf("s: %v\n", s)

}

// 重命名
func reNameFile() {
	err := os.Rename("a.txt", "b.txt")
	fmt.Printf("err: %v\n", err)
}

// 读文件
func readFile() {
	b, err := os.ReadFile("11.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("b: %v\n", string(b[:]))
	}
}

// 写文件
func writeFile() {
	s := "XF"
	os.WriteFile("11.txt", []byte(s), os.ModePerm)

}

func main() {

	createFile() //os.Create("a.txt")
	createDir()  //os.MkdirAll("test1/a/b", os.ModePerm)
	removeDir()  //os.Remove("a.txt")
	getWd()      //os.Getwd()
	chWd()       //os.Chdir("d:/")
	reNameFile() //os.Rename("a.txt", "b.txt")
	writeFile()  //os.WriteFile("11.txt", []byte(s), os.ModePerm)
	readFile()   //os.ReadFile("11.txt")
}
