package string

// ReverseVowels 翻转字符串中的元音字母, 2 2 交换 从left 和right 交换
func ReverseVowels(s string) string {
	left, right := 0, len(s)-1

	for left < right {
		if !isVowels(s[left]) {
			left++
			continue
		}
		if !isVowels(s[right]) {
			right--
			continue
		}

		// left , right 都是元音字符
		s = swap(s, left, right)
		left++
		right--
	}
	return s
}

// swap 交换字符
func swap(str string, indexA, indexB int) string {
	runes := []rune(str)
	runes[indexA], runes[indexB] = runes[indexB], runes[indexA]
	return string(runes)
}

func isVowels(inputByte byte) bool {
	input := string(inputByte)
	bucket := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	for _, val := range bucket {
		if val == input {
			return true
		}
	}
	return false
}

// removeDuplicates 移除数组中的重复数字，并且不能重新分配新数组
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	position := 0
	for i := 1; i < len(nums); i++ {
		prev := nums[position]
		current := nums[i]
		if current != prev {
			position++
			nums[position] = current
		}
	}

	return position + 1
}

// 有序数组， 最多只能重复2次
func removeDuplicatesII(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	position := 0
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		//next := nums[i+1]
		if i> 0 && i <len(nums)-1 && current == nums[i-1] && current ==  nums[i+1] {
			continue
		}
		nums[position] = current
		position++

	}
	return position
}
