package main

import "fmt"

func main() {
	// 字符串类型(String)：存文字
	var gameName string = "绝区零"

	element := "以太"

	// 整数类型(Integer)：存等级、血量
	level := 70

	// 浮点数类型(Float)：存带小数点的（伤害、暴击率）
	critRate := 75.5

	// 布尔类型(Boolean)：存对或错（是否在线）
	isOnline := true

	// 视觉化输出：不仅看值，而且看它的“盒子”是什么做的
	fmt.Println("--- 角色状态 ---")
	fmt.Printf("游戏名称：%v, 类型：%T\n", gameName, gameName)
	fmt.Printf("元素属性：%v, 类型：%T\n", element, element)
	fmt.Printf("当前等级：%v, 类型：%T\n", level, level)
	fmt.Printf("暴击率：%v, 类型：%T\n", critRate, critRate)
	fmt.Printf("是否在线：%v, 类型：%T\n", isOnline, isOnline)

	// %v：显示盒子的内容（Value）
	// %T：显示盒子的类型（Type）
}
