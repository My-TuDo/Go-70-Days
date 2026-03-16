package main

import "fmt"

func main() {
	// 1. 定义拥有的原石数量
	var myCrystals int
	fmt.Println("请输入你拥有的原石数量：")

	// 2.这里的 Scan 是一个新知识点： 它可以让程序停下来，等待用户在终端输入数字。
	fmt.Scan(&myCrystals)

	// 3. 十连抽需要 1600 原石
	price := 1600

	fmt.Println("--- 抽卡结果判定 ---")

	if myCrystals >= price {
		if myCrystals >= 16000 {
			fmt.Println("哇！大氪户，我逢魔大伟敬你！")
		}
		fmt.Println("数量足够！正在开启金光...")
		elementCheck()
	} else if myCrystals >= 160 {
		count := myCrystals / 160
		fmt.Printf("原石不足十连，但可以单抽 %v 次！\n", count)
	} else {
		// 计算还差多少个
		need := price - myCrystals
		fmt.Printf("原石不足！还差 %v 个原石才能十连抽！\n", need)
	}
}

func elementCheck() {
	element := "雷"

	switch element {
	case "火":
		fmt.Println("角色：玛薇卡 - 星级：5星！")
	case "水":
		fmt.Println("角色：哥伦比亚 - 星级：5星！")
	case "风":
		fmt.Println("角色：温迪 - 星级：5星！")
	case "雷":
		fmt.Println("角色：瓦蕾莎 - 星级：5星！")
	default:
		fmt.Println("未知的元素属性！")
	}
}
