package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//获取长度
	var Name string = "朱云飞"
	fmt.Println(len(Name))                    //获取字节长度
	fmt.Println(utf8.RuneCountInString(Name)) //获取字符长度

	//判断是否以xx开头，结尾
	fmt.Println(strings.HasPrefix(Name, "朱"))
	fmt.Println(strings.HasSuffix(Name, "飞1"))

	//是否包含特定字符
	fmt.Println(strings.Contains(Name, "云飞"))

	//更改大小写
	Str := " Mooooing "
	fmt.Println(strings.ToUpper(Str))
	fmt.Println(strings.ToLower(Str))

	//截断两头
	fmt.Println(strings.TrimRight(Str, "ing "))
	fmt.Println(strings.TrimLeft(Str, " M"))
	fmt.Println(strings.Trim(Str, " ")) //删除两边的w

	//替换
	fmt.Println(strings.Replace(Str, "o", "0", 1))  //找到o替换成0，从左到右找第一个替换
	fmt.Println(strings.Replace(Str, "o", "0", 2))  //找到o替换成0，从左到右找前两个替换
	fmt.Println(strings.Replace(Str, "o", "0", -1)) //找到o，替换所有

	//分割
	fmt.Println(strings.Split(Str, "o")) //用o来分割

	//拼接
	datalist := []string{"I am ", "Mooing"}
	fmt.Println(strings.Join(datalist, ""))
	//更高效一些，建议这种方式
	var test strings.Builder
	test.WriteString("我是")
	test.WriteString("Mooing")
	fmt.Println(test.String())

	//字符串和字节集合转换,可用来简单加密密码
	password := "123+abcXX"
	bytepwd := []byte(password)
	fmt.Println(bytepwd)

	bytelist := []byte{49, 50, 51, 43, 97, 98, 99, 88, 88}
	TargetStr := string(bytelist)
	fmt.Println(TargetStr)

	//数字和字符串的转换,可能会用在随机验证码数字和字符同时存在的情况。
	v1 := string(65)
	println(v1) //A
	v2, size := utf8.DecodeLastRuneInString("朱")
	fmt.Println(v2, size)

}
