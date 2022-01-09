package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	queue := [][2]*TreeNode{{root.Left, root.Right}}
	for len(queue) > 0 {
		l := queue[0][0]
		r := queue[0][1]
		queue = queue[1:]

		if l == nil && r == nil {
			continue
		}
		if (l == nil && r != nil) || (l != nil && r == nil) || l.Val != r.Val {
			return false
		}
		queue = append(queue, [2]*TreeNode{l.Left, r.Right})
		queue = append(queue, [2]*TreeNode{l.Right, r.Left})
	}
	return true
}
