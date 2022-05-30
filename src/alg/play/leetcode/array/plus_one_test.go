package array

import (
	"fmt"
	"testing"
)

/**
easy:

Given a non-empty array of decimal digits representing a non-negative integer,
increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list,
and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Example 3:

Input: digits = [0]
Output: [1]

Given a number represented as an array of digits, plus one to the number.
*/

func plusOne(digits []int) []int {
	if digits == nil || len(digits) < 0 {
		return []int{1}
	}
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		digits[i] = digits[i] % 10
		if digits[i] != 0 { // 末尾小于o
			return digits
		}

	}
	var result =make([]int,len(digits)+1)
	result[0]=1
	//result = append(result,digits...)
	return result
}


func Test_plusOne(t *testing.T) {
	var input = []int{1, 2, 3}
	//fmt.Println(reverseList(input))

	fmt.Println(plusOne(input))
	input = []int{4, 3, 2, 1}
	fmt.Println(plusOne(input))
	input = []int{0}
	fmt.Println(plusOne(input))
	input =[]int{9}
	fmt.Println(plusOne(input))
}
