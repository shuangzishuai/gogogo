package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := -1
	res := search(nums, target)
	fmt.Println(res)

	res2 := search2(nums, target)
	fmt.Println(res2)
}

func search(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	high := len(nums)
	low := 0
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
