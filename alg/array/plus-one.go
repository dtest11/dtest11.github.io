package array

func plusOne(digits []int) []int {
	carry := 0
	// reverse digits
	i := 0
	j := len(digits) - 1
	for i < j {
		digits[i], digits[j] = digits[j], digits[i]
		j--
		i++
	}
	var result []int
	for _, v := range digits {
		val := v + carry
		result = append(result, val%10) //注意 这里取余数，不能使用/
		carry = val / 10
	}
	if carry != 0 {
		result = append(result, carry)
	}
	i = 0
	j = len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		j--
		i++
	}
	return result
}
