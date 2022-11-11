package main

/*
接口可以通过嵌套创建新的接口，例如：飞鱼即可以飞，也可以游泳。我们可以创建一个fly和swim接口，组成飞鱼接口
*/
import "fmt"

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct{}

func (fish Fish) fly() {
	fmt.Println("fly...")
}

func (fish Fish) swim() {
	fmt.Println("swim...")
}

func main() {
	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()
}
