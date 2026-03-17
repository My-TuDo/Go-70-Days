package main

import "fmt"

// 1. 定义一个函数：计算平均战力
// 它接收一个分数切片，返回一个浮点数（平均分）
func calcAverage(scores []int) float64 {
	sum := 0
	for _, s := range scores {
		sum += s
	}
	// 因为整数除法会丢失小数部分，所以要先把sum和len(scores)强制转换成float64类型
	return float64(sum) / float64(len(scores))
}

// 魔改任务： 2. 再写一个函数，如果战力大于9500,则返回字符串“T0级别”，否则返回"主力级别"。
func checkLevel(power int) string {
	if power > 9500 {
		return "T0级别"
	}
	return "主力级别"
}

func main() {
	// 补充：map初始化问题。
	/*	如果你明确要定义多少键值并直接传入，可以使用map字面量的方式来初始化Map，就像下面powers这个变量一样。
	*	如果你不确定要存储多少数据，或者想要先创建一个空的Map再逐步添加数据，可以使用make函数来初始化Map，就像下面emptyMap这个变量一样。
	*	无论哪种方式，初始化后的Map都可以正常使用，只要你按照正确的语法来操作它就行了。
	*
	*	emptyMap := make(map[string]int) // 创建一个空的Map，键是字符串类型，值是整数类型
	*	powers := map[string]int{ // 创建一个有初始数据的Map
	*		"温迪": 8500,
	*		"钟离": 9200,
	*		"影":  7800,
	*	}
	*
	*   如果你知道大概有多少数据需要存储，可以在make函数中指定一个初始容量，这样可以提高性能，减少内存分配的次数。比如：
	*   powers := make(map[string]int, 10) // 创建一个初始容量为10的Map
	 */

	// 2. 创建一个战力表(Map)
	powers := map[string]int{
		"温迪": 8500,
		"钟离": 9200,
		"影":  7800,
	}

	// 3. 增加一条数据
	powers["纳西妲"] = 8800

	// 4. 打印所有人战力
	fmt.Println("--- 提瓦特战力板 ---")
	allScores := []int{} // 用来存储所有人的战力分数
	for name, p := range powers {
		fmt.Printf("角色： %-5s | 战力： %d\n", name, p)
		// 存入分数切片，后续计算平均战力用
		allScores = append(allScores, p)
	}

	// 5. 调用函数计算平均战力
	avg := calcAverage(allScores)
	fmt.Printf("\n团队平均战力是：%.2f\n", avg)

	// 6. 模拟查询功能
	searchName := "钟离"
	if p, exist := powers[searchName]; exist {
		fmt.Printf("\n查询结果：%s的战力是%d\n", searchName, p)
	} else {
		fmt.Printf("\n查询结果：没有找到%s\n", searchName)
	}

	// 魔改任务： 1. 删除一个角色
	delete(powers, "影")
	fmt.Println("\n--- 删除了影之后的战力板 ---")
	for name, p := range powers {
		fmt.Printf("角色： %-5s | 战力： %d\n", name, p)
	}

	// 循环遍历Map，打印出每个角色的级别
	fmt.Println("\n--- 角色级别 ---")
	for name, p := range powers {
		level := checkLevel(p)
		fmt.Printf("角色： %-5s | 战力： %d | 级别： %s\n", name, p, level)
	}

	// 补充知识点：因为go不允许定义了变量而不使用，所以我们可以用匿名变量_来占位，表示这个变量我们不需要使用它的值。
	// 比如在上面的for循环中，如果我们不需要使用角色的名字，只想要战力分数，我们可以这样写：
	/*
		for _, p := range powers {
			fmt.Printf("战力： %d\n", p)
		}
	*/
}
