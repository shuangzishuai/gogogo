package main

import "fmt"

/*
	输入: nums = [0,1,0,3,12]
	输出: [1,3,12,0,0]
*/

func main() {
	nums := []int{0, 1, 0, 3, 12}
	lens := moveZeroes(nums, 0)
	fmt.Printf("%v\n", nums[:lens])
}

func moveZeroes(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}

	return slow
}
