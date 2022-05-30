package main

type pair struct {
	Index int
	ele   int
}

func sortAnswer(data []int) []int {
	if len(data) == 2 {
		if data[0] > data[1] {
			data[0], data[1] = data[1], data[0]
		}
	}
	return data
}


func sort(array []int) []pair {
	var result []pair
	for i := 0; i <= len(array)-1; i++ {
		result = append(result, pair{
			Index: i,
			ele:   array[i],
		})
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j].ele > result[j+1].ele {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}


func twoSum(nums []int, target int) []int {
	data := sort(nums)
	for i, j := 0, len(nums)-1; i < j; {
		a := data[i].ele
		b := data[j].ele
		if a+b > target {
			j--
			continue
		} else if a+b < target {
			i++
			continue
		} else {
			res := sortAnswer([]int{data[i].Index, data[j].Index})
			return res
		}
	}
	return nil

}
