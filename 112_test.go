package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	remain := sum - root.Val
	if root.Left == nil && root.Right == nil && remain == 0 {
		return true
	}
	return hasPathSum(root.Left, remain) || hasPathSum(root.Right, remain)
}
