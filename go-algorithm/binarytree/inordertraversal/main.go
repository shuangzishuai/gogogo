package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraversal(root *TreeNode) (res []int) {
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return res
}

func main() {
	var root TreeNode
	root.Val = 1
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	fmt.Println(inOrderTraversal(&root))
}
