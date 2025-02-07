package main

import (
	"container/list"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
	Next  *treeNode
}

func connect(root *treeNode) *treeNode {
	if root == nil {
		return root
	}

	queue := list.New()
	queue.PushBack(root)
	tmpArr := make([]*treeNode, 0)
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

			tmpArr = append(tmpArr, node)
		}
		if len(tmpArr) > 1 {
			for i := 0; i < len(tmpArr)-1; i++ {
				tmpArr[i].Next = tmpArr[i+1]
			}
		}
		tmpArr = []*treeNode{}
	}

	return root

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

	connect(root)
}
