package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	queue := list.New()
	res := [][]int{}
	if root == nil {
		return res
	}
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		tmp := []int{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

func main() {
	root := &TreeNode{
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
	levelOrderBottom(root)
}
