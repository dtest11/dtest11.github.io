mod solution;

pub struct Solution {}


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

#[cfg(test)]
mod tests {
    use super::*;
    struct Arg {
        input:i32,
        expect:String,
    }

    impl Arg {
        fn new(input:i32,expect:String) -> Arg{
            Arg{input,expect}
        }
    }

    #[test]
    fn test_to_any_diget() {
        let args =vec![Arg::new(87,"1010111".to_string()),
                       Arg::new(10,"1010".to_string()),
                       Arg::new(0,"0".to_string()),
                       Arg::new(31,"11111".to_string())];
        for arg in args{
            assert_eq!(arg.expect,Solution::to_any_diget(arg.input,2));
        }
    }

    #[test]
    fn test_2_10() {
        let str_2 = "101010".to_string();
        let n = 42;
        assert_eq!(str_2, Solution::ten_two_2(42));
    }

    #[test]
    fn test_10_2() {
        let str_2 = "101010".to_string();
        let n = 42;
        assert_eq!(n, Solution::two_ten(str_2));
    }

    #[test]
    fn test_is_strictly_palindromic() {}
}

