### 进制转换
```rust

impl Solution {
    // 转换为任意进制
    pub fn to_any_diget(n: i32, target: i32) -> String {
        let mut remain = n;
        let mut result = "".to_string();
        while remain != 0 {
            let carry = remain % target;
            result.push_str(&carry.to_string());
            remain = remain / target;
        }
        if result.len() == 0 {
            return "0".to_string();
        }
        result.chars().rev().collect()
    }
    // ten_two_2 10进制转换2进制
    pub fn ten_two_2(n: i32) -> String {
        let mut result = "".to_string();
        let mut carry;
        let mut remain = n;
        while remain != 0 {
            remain = remain / 2;// 做除法
            carry = remain % 2;// 取余数 求模取余
            result.push_str(&carry.to_string());
        }
        return result;
    }

    // 2进制转换为10进制： 逆序遍历 2的0次方，2的1次方，累加
    pub fn two_ten(num: String) -> i32 {
        // reverse iter
        let mut chars: Vec<char> = num.chars().collect();
        chars.reverse();

        let mut result = 0;
        for (i, v) in chars.iter().enumerate() {
            if v.to_digit(10).unwrap() != 0 {
                result = result + (1 << i)
            }
        }
        result
    }
}


```


### 16-10
```go


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



### 分解质因数

质因数（素因数或质因子）在数论里是指能整除给定正整数的质数

```go

func fuck(n int) []int {
	var result []int
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n = n / i
				result = append(result, i)
			}
		}
	}
	if n != 1 {
		result = append(result, n)
	}
	return result
}
```

