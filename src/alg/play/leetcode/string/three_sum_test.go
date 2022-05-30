package string

//func threeSum(nums []int, target int) [][]int {
//	var record = make(map[int]int)
//	for index, v := range nums {
//		record[v] = index
//	}
//	var result []int
//	for i := 0; i < len(nums); i++ {
//		tmp := target - nums[i]
//		found := twoSumV2(nums, record, i, tmp)
//		if found != nil {
//			result = append(result, found...)
//			result = append(result, i)
//			return result
//		}
//	}
//	return nil
//}
//
//func twoSumV2(nums []int, record map[int]int, position int, target int) []int {
//	for val, index := range record {
//		need := target - val
//		if i, ok := record[need]; ok && i != index && i != position {
//			return []int{i, val}
//		}
//	}
//	return nil
//}
