package bit

import (
	"testing"
)

func divide(dividend int, divisor int) int {
	// 10 / 3 =3
	// 5/ 2 =2
	var sign int
	if oppositeSigns(dividend,divisor) {
		sign = -1
	}else {
		sign =1
	}
	if divisor < 0 {
		divisor = -divisor

	}
	if dividend < 0 {
		dividend = -dividend

	}
	if divisor > dividend {
		return 0
	}

	result := 0
	//循环条件：当被除数大于等于除数时候
	for dividend >= divisor {
		ds := divisor
		count := 1
		//<< 移位操作,向左移动1位
		///当dvd大于dvs*2的mult次方的时候,即计算dvd中最多有2的几次方个dvs
		for dividend >= ds<<1 {
			ds = ds << 1 // 2倍递增
			count = count << 1
		}
		// dividend <= ds
		//fmt.Printf("%d里面有,2的%d次方个%d\r\n", dividend, count, divisor)
		dividend = dividend - ds
		result += count
	}
	return result * sign
}

// 累计
func divideAdd(dividend int, divisor int) int {
	// 10 / 3 =3
	// 5/ 2 =2
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor > dividend {
		return 0
	}

	count := 0
	step := divisor
	for divisor <= dividend {
		divisor += step
		count++
	}
	return count
}

func oppositeSigns(x, y int) bool {
	x=x ^ y
	return x< 0
}
func Test_divide(t *testing.T) {
	//result := divideAdd(10, 3)
	//t.Log(result)
	result := divide(7, -3)
	t.Log(result)
	result = divide(7, 3)
	t.Log(result)
	t.Log(oppositeSigns(100,-100))
	t.Log(oppositeSigns(100,100))
	t.Log(oppositeSigns(10,100))
	t.Log(oppositeSigns(7,-3))
}
