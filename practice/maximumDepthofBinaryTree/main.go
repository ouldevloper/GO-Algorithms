package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth_(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode, dep int) int
	dfs = func(root *TreeNode, dep int) int {
		if root == nil {
			return dep
		}
		l := dfs(root.Left, dep+1)
		r := dfs(root.Right, dep+1)
		return max(l, r)
	}
	return dfs(root, 0)
}
func maxDepth(root *TreeNode) int {
	var dfs func(root *TreeNode, dep int) int
	dfs = func(root *TreeNode, dep int) int {
		if root == nil {
			return dep
		}
		return max(dfs(root.Left, dep+1), dfs(root.Right, dep+1))
	}
	return dfs(root, 0)
}
func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, &TreeNode{5, nil, nil}}}}
	fmt.Println("Result : ", maxDepth(root))
}
