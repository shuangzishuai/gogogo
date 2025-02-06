package main

import (
	"container/list"
	"fmt"
)

type rootNode struct {
	Val      int
	Children []*rootNode
}

func levelorder(root *rootNode) [][]int {
	queue := list.New()
	res := [][]int{}
	if root == nil {
		return res
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		var tmp []int
		for T := 0; T < length; T++ {
			// 从队列头部移除一个节点并转换为rootNode类型
			// queue.Remove(queue.Front()) 从队列中移除并返回第一个元素
			// 由于queue存储的是interface{}类型，需要用类型断言.(*rootNode)转换回rootNode指针类型
			myNode := queue.Remove(queue.Front()).(*rootNode)
			tmp = append(tmp, myNode.Val)
			for i := 0; i < len(myNode.Children); i++ {
				queue.PushBack(myNode.Children[i])
			}
		}
		res = append(res, tmp)
		fmt.Printf("res = %v", res)
	}

	return res
}

func main() {
	rootNode := &rootNode{
		Val: 1,
		Children: []*rootNode{
			{
				Val: 3,
				Children: []*rootNode{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	levelorder(rootNode)
}
