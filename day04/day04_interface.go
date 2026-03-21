package main

import (
	"fmt"
)

// 1. 定义接口：战士（只要能攻击，就是战士）
// 进阶升级：不仅能攻击，还能报名字
type Warrior interface {
	Attack() string
	GetName() string
}

// 2. 定义两个不同的结构体（职业）
type Swordman struct {
	Name string
}

type Archer struct {
	Name string
}

// 定义一个Mage结构体
type Mage struct {
	Name string
}

// 3. 让这两个职业都实现 Attack 方法
// 给结构体实现 GetName 方法
func (s Swordman) Attack() string {
	return "Swordman" + s.Name + "发动了斩击！"
}
func (s Swordman) GetName() string {
	return s.Name
}

func (a Archer) Attack() string {
	return "Archer" + a.Name + "发射了强力箭矢！"
}
func (a Archer) GetName() string {
	return a.Name
}

func (m Mage) Attack() string {
	return "Mage" + m.Name + "施放了火球术！"
}
func (m Mage) GetName() string {
	return m.Name
}

// 4. 定义一个带错误处理的函数：释放技能
// 改为fmt.Errorf动态处理错误信息
func castSkill(w Warrior, mana int) (string, error) {
	if mana < 20 {
		// 返回一个自定义错误
		return "", fmt.Errorf("%s 元素能量不足，无法释放技能！", w.GetName())
	}
	return w.Attack(), nil
}

func main041() {
	// 创建两个角色
	p1 := Swordman{Name: "神里绫华"}
	p2 := Archer{Name: "温迪"}
	p3 := Mage{Name: "可莉"}

	// 5. 统一处理： 接口的神奇之处
	// 不管是剑士还是弓箭手， 只要是Warrior，就能传给castSkill
	team := []Warrior{p1, p2, p3}

	for _, hero := range team {
		// 模拟不同能量情况
		res, err := castSkill(hero, 10) // 能量设为10，故意报错
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}

		res2, err2 := castSkill(hero, 50) // 能量足够
		if err2 == nil {
			fmt.Println(res2)
		}

	}
}
