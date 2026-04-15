package main

import (
	"errors"
	"fmt"
)

// 1. 定义一个底层的“炸弹”
var ErrDatabaseDown = errors.New("数据库物理连接断开")

// 2. 模拟中层逻辑：认证失败
func authenticate() error {
	// 包装底层错误，加上这一层的描述
	return fmt.Errorf("认证模块报错： %w", ErrDatabaseDown)
}

// 3. 模拟顶层逻辑： API 响应
func handleRequest() error {
	if err := authenticate(); err != nil {
		return fmt.Errorf("API 接口异常：%w", err)
	}
	return nil
}

func main() {
	fmt.Println("正在执行系统巡检...")

	err := handleRequest()

	if err != nil {
		// A. 直接打印，看到完整的”套娃“路径
		fmt.Println("完整报错链路：", err)

		// B. 【核心】剥开套娃，寻找真相
		// 即使经过了3层包装， IS 依然能抓到最初的 ErrDatabaseDown
		if errors.Is(err, ErrDatabaseDown) {
			fmt.Println("警报！数据库连接断开了！")
		}
	}
}
