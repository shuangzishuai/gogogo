package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// //递归实现 前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

func preOrderTraversalByStack(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		res = append(res, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return res
}

func main() {
	var root TreeNode
	root.Val = 1
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	fmt.Println(preorderTraversal(&root))        //[1 2 4 5 3]
	fmt.Println(preOrderTraversalByStack(&root)) //[1 2 4 5 3]
}
