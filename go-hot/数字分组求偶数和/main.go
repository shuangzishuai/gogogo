package main

import (
	"fmt"
	"strconv"
)

func solution(numbers []int) int {
	// Please write your code here
	// 将每个数字组转换为字符串列表
	var groups []string
	for _, num := range numbers {
		groups = append(groups, strconv.Itoa(num))
	}

	fmt.Println(groups)

	count := 0

	var generateCombinations func(index int, currentSum int)
	generateCombinations = func(index int, currentSum int) {
		// 如果已经遍历完所有数字组
		if index == len(groups) {
			// 检查当前和是否为偶数
			if currentSum%2 == 0 {
				count++
			}
			return
		}

		// 遍历当前数字组中的每个数字 递归查询
		for _, digit := range groups[index] {
			// 将字符转换为数字并累加到当前和
			num, _ := strconv.Atoi(string(digit))
			fmt.Println(num)
			generateCombinations(index+1, currentSum+num)
		}
	}

	// 从第一个数字组开始生成组合
	generateCombinations(0, 0)

	return count
}

func main() {
	//小M面对一组从 1 到 9 的数字，这些数字被分成多个小组，并从每个小组中选择一个数字组成一个新的数。目标是使得这个新数的各位数字之和为偶数。任务是计算出有多少种不同的分组和选择方法可以达到这一目标。
	// You can add more test cases here
	fmt.Println(solution([]int{123, 456, 789}) == 14)
	fmt.Println(solution([]int{123456789}) == 4)
	fmt.Println(solution([]int{14329, 7568}) == 10)
}
