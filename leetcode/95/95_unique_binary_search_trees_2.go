package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	cache := make([][][]*TreeNode, n+1)
	for i := range cache {
		cache[i] = make([][]*TreeNode, n+1)
	}
    return generateTreesHelper(1, n, cache)
}
func generateTreesHelper(start, end int, cache [][][]*TreeNode) []*TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return []*TreeNode{
			{Val: start},
		}
	}
    if cache[start][end] != nil {
		return cache[start][end]
	}
	ans := []*TreeNode{}

	for i:=start;i<=end;i++ {
		ltree := generateTreesHelper(start, i-1, cache)
		rtree := generateTreesHelper(i+1, end, cache)
		if len(ltree) == 0 {
			for _, n := range rtree {
				tmp := TreeNode{Val: i}
				tmp.Left = nil
				tmp.Right = n
				ans = append(ans, &tmp)
			}
			continue
		}
		if len(rtree) == 0 {
			for _, n := range ltree {
				tmp := TreeNode{Val: i}
				tmp.Left = n
				tmp.Right = nil
				ans = append(ans, &tmp)
			}
			continue
		}
		for _, l := range ltree {
			for _, r := range rtree {
				tmp := TreeNode{Val: i}
				tmp.Left = l
				tmp.Right = r
				ans = append(ans, &tmp)
			}
		}
	}

	cache[start][end] = ans
	return ans
}

func printTree(tarr []*TreeNode) {

	fmt.Printf("[")
	for _, v := range tarr {
		fmt.Printf("[")
		printDFS(v)
		fmt.Printf(" ]")
	}
	fmt.Printf("]")
}
func printDFS(t *TreeNode) {
	if t == nil {
		return
	}
	printDFS(t.Left)
	printDFS(t.Right)
	fmt.Printf(" %d", t.Val)
}

func main() {
	printTree(generateTrees(3))
}