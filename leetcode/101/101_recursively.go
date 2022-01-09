package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return areSymmetric(root.Left, root.Right)
}

func areSymmetric(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if (l == nil && r != nil) || (l != nil && r == nil) || l.Val != r.Val {
		return false
	}
	return areSymmetric(l.Left, r.Right) && areSymmetric(l.Right, r.Left)
}
