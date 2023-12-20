package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	// res := removeElement(nums, val)
	// fmt.Println(res)

	res1 := removeElement1(nums, val)
	fmt.Println(res1)
}

// 双向指针法
func removeElement1(nums []int, val int) int {
	length := len(nums)
	res := 0                      //慢指针
	for i := 0; i < length; i++ { //快指针
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	fmt.Println(nums)
	return res
}

// 相向双指针
func removeElement(nums []int, val int) int {
	// 有点像二分查找的左闭右闭区间 所以下面是<=
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		//各自找到后开始覆盖 覆盖后继续寻找
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	fmt.Println(nums)
	return left
}
