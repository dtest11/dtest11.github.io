package main

/***
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Notice that the solution set must not contain duplicate quadruplets.
=


Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [], target = 0
Output: []

*/

//type pair struct {
//	Index int
//	ele   int
//}
//
//func sortAnswer(data []int) []int {
//	if len(data) == 2 {
//		if data[0] > data[1] {
//			data[0], data[1] = data[1], data[0]
//		}
//	}
//	return data
//}
//
//
//func sort(array []int) []pair {
//	var result []pair
//	for i := 0; i <= len(array)-1; i++ {
//		result = append(result, pair{
//			Index: i,
//			ele:   array[i],
//		})
//	}
//
//	for i := 0; i < len(result); i++ {
//		for j := 0; j < len(result)-1-i; j++ {
//			if result[j].ele > result[j+1].ele {
//				result[j], result[j+1] = result[j+1], result[j]
//			}
//		}
//	}
//	return result
//}

func fourSum(nums []int, target int) [][]int {
	return fourSumTarget(nums, target)
}

func fourSumTarget(nums []int, target int) [][]int {
	var result [][]int
	data := sort(nums)
	length := len(data)
	for i := 0; i < len(data)-3; i++ {
		if i > 0 && data[i].ele == data[i-1].ele {
			continue // 去重
		}
		if data[i].ele+data[i+1].ele+data[i+2].ele+data[i+3].ele > target {
			break
		}
		if data[i].ele+data[length-3].ele+data[length-2].ele+data[length-1].ele < target {
			continue
		}
		current := data[i]
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && data[j].ele == data[j-1].ele {
				continue
			}
			if current.ele+data[j].ele+data[j+1].ele+data[j+2].ele > target {
				break
			}
			if current.ele+data[length-2].ele+data[length-1].ele < target {
				continue
			}
			left := j + 1 // three sum 的过程
			right := length - 1
			current2 := data[j]
			for left < right {
				sum := current2.ele + current.ele + data[left].ele + data[right].ele
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					result = append(result, []int{current.ele, current2.ele, data[left].ele, data[right].ele})
					// 去重
					for left < right && data[left].ele == data[left+1].ele {
						left++
					}
					// 去重
					for left < right && data[right].ele == data[right-1].ele {
						right--
					}
					left++
					right--
				}
			}
		}
	}

	return result
}

func fourSum1(nums []int, target int) (quadruplets [][]int) {
	//sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/4sum/solution/si-shu-zhi-he-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
