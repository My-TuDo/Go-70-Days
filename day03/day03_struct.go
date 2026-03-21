package main

import "fmt"

// 1. 定义一个结构体：游戏角色
type Player struct {
	Name   string
	Level  int
	Health int
	// 新增字段：武器
	Weapon struct {
		Name   string
		Damage int
	}
}

// 2. 定义一个函数：升级（*代表接受指针）
// 如果不加*，升级只会在函数内部发生，外部角色等级不变。
func levelUp(p *Player) {
	p.Level++
	fmt.Printf("%s 升级了！当前等级：%d\n", p.Name, p.Level)
}

// 写一个方法：装备武器
func (p *Player) equipWeapon(name string, damage int) {
	p.Weapon.Name = name
	p.Weapon.Damage = damage
	fmt.Printf("%s 装备了 %s，伤害：%d\n", p.Name, p.Weapon.Name, p.Weapon.Damage)
}

func main0001() {
	// 3. 创建一个角色实例
	// 使用 & 取地址。创建的是一个指针对象
	p1 := &Player{
		Name:   "旅行者",
		Level:  1,
		Health: 100,
	}

	fmt.Printf("初始状态：%+v\n", p1)

	// 4. 调用函数修改原件
	levelUp(p1)
	levelUp(p1)

	fmt.Printf("最终状态：%+v\n", p1)

	// 视觉化理解指针
	fmt.Println("--------------------------------")
	fmt.Printf("变量p1存储的地址：%p\n", p1)

	// 视觉派重点： %+v 和%p
	// %+v 会打印结构体的字段名和对应的值，方便我们查看结构体的内容。
	// %p 会打印变量存储的内存地址，帮助我们理解指针的概念。

	// 5. 调用方法装备武器
	p1.equipWeapon("无锋剑", 150)

	fmt.Printf("最终状态：%+v\n", p1)
}
