/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dequeue(q *[]TreeNode) *TreeNode {
	if len(*q) <= 0 {
		return nil
	}
	ret := (*q)[0]
	*q = ((*q)[1:])
	return &ret
}

func enqueue(t *TreeNode, q *[]TreeNode) {
	if t == nil {
		return
	}
	*q = append(*q, *t)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pQueue := []TreeNode{}
	qQueue := []TreeNode{}
	enqueue(p, &pQueue)
	enqueue(q, &qQueue)
	for p, q = dequeue(&pQueue), dequeue(&qQueue); p != nil && q != nil; p, q = dequeue(&pQueue), dequeue(&qQueue) {
		if p.Val != q.Val {
			return false
		}
		enqueue(p.Left, &pQueue)
		enqueue(q.Left, &qQueue)
		if len(pQueue) != len(qQueue) {
			return false
		}
		enqueue(p.Right, &pQueue)
		enqueue(q.Right, &qQueue)
		if len(pQueue) != len(qQueue) {
			return false
		}
	}
	return p == nil && q == nil
}
