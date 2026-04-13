package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string { return "致命错误！！" }

func getError(isFine bool) error {
	var err *MyError = nil
	if isFine {
		return err
	}
	return nil
}

func main() {
	fmt.Println("正在进行接口空值实验...")

	err := getError(true)

	if err != nil {
		fmt.Printf("警报！捕获到异常，类型是: %T, 值是: %v\n", err, err)
		fmt.Println("为什么明明没有报错，逻辑却进入了if err != nil?")
	} else {
		fmt.Println("系统正常")
	}
}
