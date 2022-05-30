package main

import "fmt"

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for index, v := range nums {
		pre += v
		if v, ok := m[pre-k]; ok {
			count += v
			fmt.Printf("found %d %d %d %v\n", index, v, pre-k, m)
		}
		m[pre]++
	}
	return count
}

// func main() {
// 	nums := []int{1, 2, 3, 4}
// 	target := 4
// 	fmt.Println(subarraySum(nums, target))
// }
