package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// 1. 准备要发送的数据（GO结构体
	data := map[string]string{
		"act_id": "e202311201442471",
	}

	// 2. 序列化为JSON字节流
	jsonData, _ := json.Marshal(data)

	// 3. 构造POST请求
	// bytes.NewBuffer 将字节数组包装成“流”， 因为http要求读取流
	url := "http://httpbin.org/post" // 一个专门用来测试的网站
	req, _ := http.NewRequest("PSOT", url, bytes.NewBuffer(jsonData))

	// 4. 设置必要的 Header
	req.Header.Set("Content-Type", "application/json")

	// 5. 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("POST 请求发送成功，状态码:", resp.StatusCode)
}
