package main

import "math"

// Not optimal but O(n)
func getMinimumDifference(root *TreeNode) int {
	var sortedList []int
	traverse(root, &sortedList)
	diff := math.MaxInt32
	prev := sortedList[0]
	for i := 1; i < len(sortedList); i++ {
		var localDiff int
		localDiff = sortedList[i] - prev
		if diff > localDiff {
			diff = localDiff
		}
		prev = sortedList[i]
	}
	return diff
}

func traverse(root *TreeNode, sortedList *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, sortedList)
	*sortedList = append(*sortedList, root.Val)
	traverse(root.Right, sortedList)
}
