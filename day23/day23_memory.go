package main

import "fmt"

type Node struct {
	Val int
}

func main231() {
	// a 在栈上
	// 直接定义的变量通常在栈上分配，除非它们逃逸到堆上。
	a := 10

	// n 在堆上（因为用了指针，且可能逃逸）
	// 通过结构体指针分配的变量通常在堆上分配，尤其是当它们被返回或传递到其他函数时。
	n := &Node{Val: 100}

	fmt.Printf("栈变量 a 的地址： %p\n", &a)
	fmt.Printf("堆变量 n 的地址： %p\n", n)
}
