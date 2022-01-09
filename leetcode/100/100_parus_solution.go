// https://leetcode.com/problems/same-tree/discuss/1617832/Go-Level-traversal
func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := [][2]*TreeNode{{p, q}}
	for len(queue) > 0 {
		p = queue[0][0]
		q = queue[0][1]
		queue = queue[1:]

		if p == nil && q == nil {
			continue
		}
		if (p == nil && q != nil) ||
			(p != nil) && (q == nil) ||
			(p.Val != q.Val) {
			return false
		}
		queue = append(queue, [2]*TreeNode{
			p.Left,
			q.Left,
		})
		queue = append(queue, [2]*TreeNode{
			p.Right,
			q.Right,
		})
	}
	return true
}
