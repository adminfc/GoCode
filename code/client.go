package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	listener, err := net.Listen("tcp", "10.17.92.20:2701")
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	defer listener.Close()
	fmt.Println("监听成功...")

	var connectionsLock sync.Mutex
	connections := make(map[int]net.Conn)
	id := 1
	currentConn := make(chan net.Conn)

	go func() {
		for {
			conn, conn_err := listener.Accept()
			if conn_err != nil {
				fmt.Println("服务端连接失败:", conn_err)
				return
			}

			connectionsLock.Lock()
			connections[id] = conn
			connectionsLock.Unlock()

			// Notify the main goroutine about the new connection
			currentConn <- conn
			id++
		}
	}()

	for {
		conn := <-currentConn
		go handleConnection(conn, id, connections, currentConn)
	}
}

func handleConnection(conn net.Conn, id int, connections map[int]net.Conn, currentConn chan net.Conn) {
	defer conn.Close()
	fmt.Printf("有主机上线 %d ，%v\n", id, conn.RemoteAddr())

	for {
		handleCMD(id, conn, connections, currentConn)
	}
}

func handleCMD(id int, conn net.Conn, connections map[int]net.Conn, currentConn chan net.Conn) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("当前连接，%v,输入命令或输入 'L' 列出连接列表：\n", conn.RemoteAddr().String())
	input, _ := reader.ReadString('\n')

	if strings.TrimSpace(input) == "L" {
		// 列出连接列表
		fmt.Println("连接列表：")
		for i := 1; i <= len(connections); i++ {
			fmt.Printf("%d,%v, %d\n", i, conn.RemoteAddr().String(), i)
		}

		// 接受用户输入以切换连接
		var switchTo int
		fmt.Print("输入要切换到的连接序列号（或输入 0 以保留当前连接）：")
		_, err := fmt.Scanf("%d", &switchTo)

		if err != nil {
			fmt.Println("无效输入")
			//return
		}
		fmt.Printf("switchTo: %v\n", switchTo)
		if switchTo != 0 {
			if connToSwitch, ok := connections[switchTo]; ok {
				currentConn <- connToSwitch // Notify main goroutine to switch to the selected connection
				return
			}
		}

	} else {
		// 处理普通命令

		fmt.Printf("input: %v\n", len(input))
		if len(input) < 2 {
			input = "echo '输入为空'"
		} else {

			fmt.Printf("xx: %v\n", input)
			n, n_err := conn.Write([]byte(input)) //将输入的命令转换成字节，并发送到
			if n_err != nil {
				fmt.Println("发送命令失败:", n_err)

				return
			}
			fmt.Printf("发送命令 %v\n", input)

			// 接收服务端的响应
			response := make([]byte, 2048)
			rspbuf := bufio.NewReader(conn)
			n, n_err = rspbuf.Read(response)
			if n_err != nil {
				fmt.Println("接收响应失败:", n_err)
				return
			}
			rsp := string(response[:n])

			fmt.Println(rsp)
		}

	}

}
