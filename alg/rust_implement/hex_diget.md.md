```go
# 16-10

package main

import "fmt"

func hexDigitToInt(hexDigit byte) int {
	switch hexDigit {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'a', 'A':
		return 10
	case 'b', 'B':
		return 11
	case 'c', 'C':
		return 12
	case 'd', 'D':
		return 13
	case 'e', 'E':
		return 14
	case 'f', 'F':
		return 15
	default:
		return 0
	}
}

func intPow(base, exp int) int {
	result := 1

	for i := 0; i < exp; i++ {
		result *= base
	}

	return result
}

func hexToDec(hexStr string) int {
	var dec, pow int

	for i := len(hexStr) - 1; i >= 0; i-- {
		digit := hexDigitToInt(hexStr[i])
		dec += digit * intPow(16, pow)
		pow++
	}

	return dec
}
func main() {
	var input string
	_, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		panic(err)
	}
	fmt.Println(hexToDec(input))
}

```