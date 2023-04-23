package dynamic_programming

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Test_fib(t *testing.T) {
	n := 10
	t.Log(fib(n))
	t.Log(fib2(n))
}

func Test_maxSubArray(t *testing.T) {
	var a = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray(a))
}

func Test_longest_increasing_subsequence(t *testing.T) {
	var nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := lengthOfLIS(nums)
	t.Log(result)
}

func Test_longestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"
	res := longestCommonSubsequence(text1, text2)
	assert.Equal(t, 3, res)
}

func Test_lengthOfLIS_binary_search(t *testing.T) {
	var nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLISV2(nums)
	assert.Equal(t, 4, res)

	res = lengthOfLIS_binary_search(nums)
	assert.Equal(t, 4, res)
}

func Test_maxProfit_k_Two(t *testing.T) {
	t.Log(maxProfit_N(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func Test_rob(t *testing.T) {
	var nums = []int{1, 2, 3, 1}
	t.Log(rob(nums))
	t.Log(rob([]int{2, 1, 1, 2}))
	t.Log(robCircle([]int{1, 2, 3, 1}))
	t.Log(robCircle([]int{2, 3, 2}))
	t.Log(robCircle([]int{1, 2, 3}))
}

func Test_longestPalindrome(t *testing.T) {
	//s := "babad"
	//assert.Equal(t, "aba", longestPalindrome(s))
	fmt.Println(longestPalindrome("cbbd"))
}

func Test_hello(t *testing.T) {
	var input string
	fmt.Scanf("%s\n", &input)
	//var input = `A10;S20;W10;D30;X;A1A;B10A11;;A10;`
	data := strings.Split(input, ";")
	var result Point
	for _, v := range data {
		point := parsePoint(v)
		if point.X != 0 {
			result.X = result.X + point.X
		} else if point.Y != 0 {
			result.Y = result.Y + point.Y
		} else {
			continue
		}
	}
	fmt.Println(result.X, result.Y)
}

type Point struct {
	X int
	Y int
}

func parsePoint(str string) Point {
	if len(str) == 0 {
		return Point{}
	}
	symbol := str[0]
	var point Point

	switch symbol {
	case 'A':
		point.X = -1
	case 'D':
		point.X = 1
	case 'W':
		point.Y = 1
	case 'S':
		point.Y = -1
	default:
		return Point{}
	}
	var temp string
	for i := 1; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			temp += string(str[i])
			continue
		}
		return Point{}
	}
	num, _ := strconv.Atoi(temp)
	if point.X != 0 {
		return Point{X: num * point.X}
	}
	return Point{Y: num * point.Y}
}
