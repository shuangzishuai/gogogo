package levelordertraversal

//二叉树层序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelordertraversal(root *TreeNode) [][]int {
	arr := [][]int{}
	depth := 0

	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)
	return arr
}

func main() {
	//队列实现
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
	levelordertraversal(root)
}
