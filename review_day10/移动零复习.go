package main

func moveZeroes(nums []int) {
	// 定义一个慢指针，指向下一个非零元素应该放置的位置
	slow := 0

	// 定义一个快指针遍历切片，寻找非零元素
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
