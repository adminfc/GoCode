package main

import "fmt"

/*
接口和类型的关系：
1。一个类型可以实现多个接口
2。多个类型可以实现同一个接口（多态）

示例一：
比如一个Player接口，可以播放音乐，有一个Video接口可以播放视频，一个手机Mobile实现这两个接口，即可音乐也可视频。
*/

//定义一个Player接口
type Player interface {
	playMusic()
}

//定义一个Video接口
type Video interface {
	playVideo()
}

//定义Mobile结构体
type Mobile struct {
}

//实现两个接口
func (m Mobile) playMusic() {
	fmt.Println("play music...")
}
func (m Mobile) playVideo() {
	fmt.Println("play video...")
}

//示例二,多个类型实现同一个接口，比如一个宠物接口，猫狗都可以实现该接口，都可以把狗猫当作宠物类型对待，在其他语言中叫多态

type Pet interface {
	eat()
}
type Dog struct{}
type Cat struct{}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func main() {
	m := Mobile{}
	m.playMusic()
	m.playVideo()
	fmt.Println("-----------示例分割线-------------")

	var pet Pet
	pet = Dog{}
	pet.eat()
	pet = Cat{}
	pet.eat()

}
