package main

func removeDuplicates(nums []int) int {
	// 先判断输入是否为空，空则返回0
	if len(nums) == 0 {
		return 0
	}

	// 定义一个慢指针，指向下一个不重复元素应该放置的位置
	slow := 0

	// 定义一个快指针遍历切片，寻找不重复元素
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	// 返回不重复元素的数量
	return slow + 1
}
