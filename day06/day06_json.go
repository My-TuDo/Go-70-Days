package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 1. 定义用户结构体
type UserProfile struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	IsVip     bool      `json:"is_vip"`
	CreatedAt time.Time `json:"created_at"` // 创建记录时间
	// 魔改任务
	Tags []string `json:"tags"`
}

// 魔改任务：写一个函数，计算用户“创建时间”距离现在已经过去多少秒
func (u *UserProfile) SecondsSinceCreation() int64 {
	return int64(time.Since(u.CreatedAt).Seconds())
}

func main() {
	// --- 第一部分： GO结构体转JSON(发送给前端)字符串 ---
	user := UserProfile{
		Username:  "Gopher旅行者",
		Email:     "gopher@mihoyo.com",
		Age:       21,
		IsVip:     true,
		CreatedAt: time.Now(),
		Tags:      []string{"二次元", "探索", "旅行"},
	}

	// 2. 序列化（打包）
	// MarshalIndent 可以让生成的JSON带缩进，方便肉眼观察
	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("--- 序列化结果（JSON） ---")
	fmt.Println(string(jsonData))

	// --- 第二部分： JSON转GO结构体（解析前端传来的数据）---
	rawJson := `{"username": "米哈游新玩家", "age": 22, "is_vip": false}`

	var newUser UserProfile
	// 3. 反序列化（解包）
	// 注意： 这里必须传 &newUser（指针），因为函数要修改它的值
	err := json.Unmarshal([]byte(rawJson), &newUser)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return
	}

	fmt.Println("--- 反序列化结果（GO结构体） ---")
	fmt.Printf("用户名： %s, 年龄： %d, 是否VIP： %v\n", newUser.Username, newUser.Age, newUser.IsVip)

	// --- 第三部分： 时间处理（大厂高频考点）---
	fmt.Println("\n--- 处理时间 ---")
	now := time.Now()
	// 格式化时间的“奇葩”口诀： 2006-01-02 15:04:05 （记忆： 123456）
	fmt.Println("当前时间：", now.Format("2006-01-02 15:04:05"))

	// 计算用户创建时间距离现在已经过去多少秒
	time.Sleep(2 * time.Second)

	seconds := user.SecondsSinceCreation()
	fmt.Printf("用户 %s 创建时间距离现在已经过去 %d 秒\n", user.Username, seconds)
}
