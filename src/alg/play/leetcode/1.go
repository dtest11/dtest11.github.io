package leetcode

import "fmt"

/****

 字典 key 是值
 value 是索引
***/
func twoSum(nums []int, target int) []int {
	dict := make(map[int][]int)
	for index, val := range nums {

		if n, ok := dict[val]; ok {
			dict[val] = append(n, val)
		} else {
			dict[val] = []int{index}
		}
		fmt.Println(val,index,dict)

	}
	var result []int
	for i, v := range nums {
		miner := target - v
		if n, ok := dict[miner]; ok && len(n) != 1 && n[0] != i { // 排除自己
			return append(n, i)
		}
	}	
	return result
}
