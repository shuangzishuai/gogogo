package main

import (
	"container/list"
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largeValue(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	ans := make([]int, 0)
	temp := math.MinInt64

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			temp = max(temp, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}
		ans = append(ans, temp)
		temp = math.MinInt64
	}
	return ans
}

func main() {
	rootNode := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	largeVal := largeValue(rootNode)
	fmt.Printf("large value = %v", largeVal)
}
