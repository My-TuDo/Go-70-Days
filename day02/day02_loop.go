package main

import "fmt"

func main() {
	// 1. 定义一个切片：初始武器库
	weapons := []string{"无锋剑", "和璞鸢", "西风长枪", "祭礼残章", "匣里龙吟"}

	// 2. 增加一个新武器
	weapons = append(weapons, "雾切之回光")

	fmt.Println("--- 我的武器库 ---")

	// 3. 使用for range 循环遍历
	// index 是序号(0, 1, 2...), name是具体的武器名字
	for i, name := range weapons {
		fmt.Printf("[%d]装备名称：%s\n", i+1, name)
	}

	// 4. 进阶： 循环中嵌套判断
	fmt.Println("\n--- 正在搜索包含‘西风’的武器 ---")
	for _, name := range weapons {
		// 如果名字厘米带西风，就输出它
		if name == "西风长枪" {
			fmt.Printf("找到了！%s \n", name)
		}
	}

	// 魔改任务： 创建一个整数切片，存五个数并用for计算总和。
	scores := []int{85, 90, 78, 92, 88}
	nums := 0
	for _, score := range scores {
		nums += score
	}
	fmt.Printf("\n总和是：%d\n", nums)
}
