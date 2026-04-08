package main

import (
	"fmt"
	"reflect"
)

// 反射：在运行时动态获取类型信息和操作对象的能力

// 定义一个监控配置结构体
type MonitorConfig struct {
	IP   string `sre:"host"` // 自定义标签sre
	Port int    `sre:"port"`
	ID   int    `sre:"-"`
}

func inspectStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// 确保传进来的是结构体
	if t.Kind() != reflect.Struct {
		fmt.Println("错误：只能检查结构体")
		return
	}

	fmt.Printf("正在检查结构体：%s, 共有 %d 个字段\n", t.Name(), t.NumField())

	// 遍历所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // 获取字段的类型信息
		value := v.Field(i) // 获取字段的值

		// 获取自定义的sre标签
		tag := field.Tag.Get("sre")
		if tag == "-" {
			continue // 跳过被标记为 "-" 的字段
		}
		fmt.Printf("字段名: %-5s | 类型: %-8v | 值: %-10v | sre标签: %s\n", field.Name, field.Type, value, tag)
	}
}

func main() {
	conf := MonitorConfig{
		IP:   "127.0.0.1",
		Port: 3306,
		ID:   999,
	}

	fmt.Println("=== 反射检查 MonitorConfig ===")
	inspectStruct(conf)
}
