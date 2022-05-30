package geekbang

import (
	"fmt"
	"strconv"
	"strings"
)

/***
先进去 后出来

***/

type stack struct {
	top   int
	datas []int
}

// IsMatch ---
func IsMatch(input string) bool {
	inputs := strings.Split(input, "")
	choices := make(map[string]string)
	choices[")"] = "("
	choices["}"] = "{"
	var temp []string
	for _, val := range inputs {
		if _, err := strconv.Atoi(val); err != nil {
			if val == "(" || val == "{" {
				temp = append(temp, val)
			}
			if val == ")" || val == "}" {
				// 弹出最上层的元素
				if len(temp) != 0 {
					top := temp[len(temp)-1]
					expect, ok := choices[val]
					if ok && top == expect {
						temp = temp[0 : len(temp)-1]
					} else {
						return false
					}
				}

			}
		}
	}

	if len(temp) != 0 {
		return false
	}
	return true
}

func main() {
	input := `{(1*2)}`
	result := IsMatch(input)
	fmt.Println(result)
}
