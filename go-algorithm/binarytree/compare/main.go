package main

import "fmt"

//对称二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func defs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return defs(left.Left, right.Right) && defs(right.Left, left.Right)
}

func isSymmetric(root *TreeNode) bool {
	return defs(root.Left, root.Right)
}

func main() {
	//对称二叉树
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	res := isSymmetric(root)
	fmt.Printf("res = %v\n", res)
}
