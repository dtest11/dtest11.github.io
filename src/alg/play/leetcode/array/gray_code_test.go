package array

import (
	"fmt"
	"testing"
)

/***
https://www.cnblogs.com/zhuruibi/p/8988044.html#:~:text=%E6%A0%BC%E9%9B%B7%E7%A0%81%E7%AE%80%E4%BB%8B,%E7%A0%81%E3%80%81%E6%9C%80%E5%B0%8F%E5%B7%AE%E9%94%99%E7%A0%81%E7%AD%89%E3%80%82
***/

/**
生成二进制格雷码方式1：以二进制为0值的格雷码为第零项，第一项改变最右边的位元，第二项改变右起第一个为1的位元的左边位元，第三、四项方法同第一、二项，如此反复，即可排列出n个位元的格雷码。

假设原始的值从0开始，格雷码产生的规律是：
第一步，改变最右边的位元值；
第二步，改变右起第一个为1的位元的左边位元；
第三步，第四步重复第一步和第二步，直到所有的格雷码产生完毕（换句话说，已经走了(2^n) - 1 步）。
用一个例子来说明：
　　假设产生3位元的格雷码，原始值位 000
　　第一步：改变最右边的位元值： 001
　　第二步：改变右起第一个为1的位元的左边位元： 011
　　第三步：改变最右边的位元值： 010
　　第四步：改变右起第一个为1的位元的左边位元： 110
　　第五步：改变最右边的位元值： 111
　　第六步：改变右起第一个为1的位元的左边位元： 101
　　第七步：改变最右边的位元值： 100
**/

func grayCode(n int) []int {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []int{0, 1}
	}
	var result [][]int
	var data []int
	var base = make([]int, n)
	result = append(result, base)
	data = append(data,Str2DEC(base))
	right := len(base) - 1
	shouldContinue := true
	for shouldContinue {
		base[right] = zeroOne(base[right])
		data = append(data,Str2DEC(base))

		result = append(result, base)

		for r := len(base) - 1; r >= 0; r-- {
			cur := base[r]
			shouldContinue = cur == 1 && r-1 >= 0
			if cur == 1 && r-1 >= 0 {
				base[r-1] = zeroOne(base[r-1])
				data = append(data,Str2DEC(base))
				result = append(result, base)
				break
			}
		}
	}

	//for _, v := range result {
	//	fmt.Println(v)
	//}
	//for _,v := range data{
	//	fmt.Println(v)
	//}

	return  data
}

func zeroOne(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}

func Str2DEC(s []int) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}

func Test_grayCode(t *testing.T) {
	fmt.Println(grayCode(2))
}
