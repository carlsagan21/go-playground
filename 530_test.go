package main

import (
	"github.com/carlsagan21/go-playground/lib"
	"math"
)

// Not optimal but O(n)
func getMinimumDifference(root *lib.TreeNode) int {
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

func traverse(root *lib.TreeNode, sortedList *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, sortedList)
	*sortedList = append(*sortedList, root.Val)
	traverse(root.Right, sortedList)
}
