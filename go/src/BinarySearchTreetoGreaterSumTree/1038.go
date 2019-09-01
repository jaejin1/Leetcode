package BinarySearchTreetoGreaterSumTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	search(root, 0)
	return root
}

func search(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}

	root.Val += search(root.Right, sum)
	return search(root.Left, root.Val)
}
