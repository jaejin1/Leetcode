package MaximumDepthofBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ltree := maxDepth(root.Left)
	rtree := maxDepth(root.Right)

	if ltree > rtree {
		return ltree + 1
	} else {
		return rtree + 1
	}
}
