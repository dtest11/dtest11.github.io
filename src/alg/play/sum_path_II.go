package main

import (
	"github.com/9999-dollor/algorithm/define"
)

func pathSum(root *define.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	hashMap := map[int]int{0: 1}
	return pathSumHelp(root, hashMap, 0, sum)
}

func pathSumHelp(root *define.TreeNode, hashMap map[int]int, currentSum int, target int) int {
	if root == nil {
		return 0
	}
	var result int
	currentSum += root.Val
	existVal := hashMap[currentSum-target]
	result += existVal
	hashMap[currentSum] = 1
	result += pathSumHelp(root.Left, hashMap, currentSum, target)
	result += pathSumHelp(root.Right, hashMap, currentSum, target)

	hashMap[currentSum] = hashMap[currentSum] - 1
	return result
}
