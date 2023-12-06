package main

//test
import (
	"fmt"
	"io"
	"os"
)

/*
一、文件读操作：
*/

// 打开关闭文件
func openCloseFile() {
	f, _ := os.Open("11.txt") //只读
	fmt.Printf("f.Name(): %v\n", f.Name())

	f2, _ := os.OpenFile("11.txt", os.O_RDWR|os.O_CREATE, 755) //如果文件不存在，就创建它
	fmt.Printf("f2.Name(): %v\n", f2.Name())

	err := f.Close()
	fmt.Printf("err: %v\n", err)
	err2 := f2.Close()
	fmt.Printf("err2: %v\n", err2)

}

// 创建文件
func createFile() {
	f, _ := os.Create("12.txt") //等价于:openfile(name,o_rdwr|o_crate|o_trunc,0666)
	fmt.Printf("f.Name(): %v\n", f.Name())

	f2, _ := os.CreateTemp("", "temp") //第一个参数，目录默认，Temp第二个参数，是文件名前缀
	fmt.Printf("f.Name(): %v\n", f2.Name())
}

//读操作

func readOps() {
	f, _ := os.Open("11.txt")
	for { //这里之所以用循环，是考虑文件比较大的情况，可以按指定的字节长度一点点的读
		buf := make([]byte, 10)
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
	}
	f.Close()
}

//从指定位置开始读

func ReadAtFile() {
	buf := make([]byte, 10)
	f, _ := os.Open("11.txt")
	n, _ := f.ReadAt(buf, 5) //从5开始读10字节
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()

}

func seekFile() {
	f, _ := os.Open("11.txt")
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
}

func main() {
	//openCloseFile()
	//createFile()
	//readOps()
	//ReadAtFile()
	seekFile()
}
