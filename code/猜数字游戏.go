package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 使用当前时间作为随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成一个1到100之间的随机数
	secretNumber := rand.Intn(100) + 1

	fmt.Println("欢迎来到猜数字游戏！")
	fmt.Println("我已经选好了一个1到100之间的秘密数字。")

	for {
		var guess int
		fmt.Print("请猜猜是多少：")
		fmt.Scan(&guess)

		// 使用if语句判断猜测的数字与秘密数字的关系
		if guess < secretNumber {
			fmt.Println("猜小了！再试一次。")
			// 使用continue语句继续下一轮循环
			continue
		} else if guess > secretNumber {
			fmt.Println("猜大了！再试一次。")
			// 使用continue语句继续下一轮循环
			continue
		} else {
			fmt.Println("恭喜你，猜对了！")
			// 使用break语句退出循环
			break
		}
	}

	// 使用goto语句展示游戏结束消息
	fmt.Println("游戏结束。谢谢参与！")

	// 使用goto语句跳转到游戏结束标签
	goto end

end:
	fmt.Println("这是游戏结束的标签。")
}
