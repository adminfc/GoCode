package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	url := "https://example.com/api/endpoint" // 替换成您要发送POST请求的URL

	// 准备POST请求的数据
	payload := []byte(`{"key1": "value1", "key2": "value2"}`)

	// 发送POST请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("发送POST请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("请求失败，状态码: %d\n", resp.StatusCode)
		return
	}

	// 读取响应数据
	var responseBytes []byte
	_, err = resp.Body.Read(responseBytes)
	if err != nil {
		fmt.Println("读取响应数据失败:", err)
		return
	}

	// 打印响应数据
	fmt.Println("响应数据:", string(responseBytes))
}
