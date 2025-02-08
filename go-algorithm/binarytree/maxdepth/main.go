package main

import (
	"container/list"
	"fmt"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func maxDeepth(root *treeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*treeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans++
	}
	return ans
}

func main() {
	rootNode := &treeNode{
		Val: 1,
		Left: &treeNode{
			Val: 2,
			Left: &treeNode{
				Val: 4,
			},
			Right: &treeNode{
				Val: 5,
			},
		},
		Right: &treeNode{
			Val: 3,
			Left: &treeNode{
				Val: 6,
			},
			Right: &treeNode{
				Val: 7,
			},
		},
	}
	ans := maxDeepth(rootNode)
	fmt.Printf("max deepth = %v", ans)
}
