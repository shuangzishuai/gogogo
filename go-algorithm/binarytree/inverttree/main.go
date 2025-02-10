package main

import (
	"fmt"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func invertTreee(root *treeNode) *treeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTreee(root.Left)
	invertTreee(root.Right)
	return root
}

func traversalTree(root *treeNode) []int {
	var traversal func(node *treeNode)
	res := []int{}
	traversal = func(node *treeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

func main() {
	root := &treeNode{
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
	orgRoot := traversalTree(root)
	fmt.Printf("original tree = %v", orgRoot)
	invertRes := invertTreee(root)
	// fmt.Printf("invert result = %v", invertRes)
	invertRoot := traversalTree(invertRes)
	fmt.Printf("invert tree = %v", invertRoot)
}
