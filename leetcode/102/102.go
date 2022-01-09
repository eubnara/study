package leetcode

func levelOrder(root *TreeNode) [][]int {
    ans := [][]int{}
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        localAns := []int{}
        newQueue := []*TreeNode{}
        for len(queue) > 0 {
            node := queue[0]
            queue = queue[1:]
            if node == nil {
                continue
            }
            localAns = append(localAns, node.Val)
            newQueue = append(newQueue, node.Left, node.Right)
        }
        if len(localAns) > 0 {
            ans = append(ans, localAns)
        }
        queue = newQueue
    }
    return ans
}
