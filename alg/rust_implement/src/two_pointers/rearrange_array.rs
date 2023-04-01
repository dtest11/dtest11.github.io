struct Solution;
impl Solution {
    pub fn pivot_array(nums: Vec<i32>, pivot: i32) -> Vec<i32> {
        let mut n1: Vec<i32> = Vec::new();
        let mut n2: Vec<i32> = Vec::new();
        let mut n3: Vec<i32> = Vec::new();
        for v in nums {
            if v < pivot {
                n1.push(v);
            } else if v == pivot {
                n2.push(v);
            } else {
                n3.push(v);
            }
        }
        n1.append(&mut n2);
        n1.append(&mut n3);
        n1
    }

    // 重排数组，形成【正数，负数，正数，负数]的一对一对结构
    // https://leetcode.com/problems/rearrange-array-elements-by-sign/description/
    pub fn rearrange_array(nums: Vec<i32>) -> Vec<i32> {
        let mut result: Vec<i32> = Vec::new();
        let mut a = 0;
        let mut b = 0;
        let max = nums.len();
        while result.len() < max {
            // 找到第一个正数
            for i in a..max {
                if nums[i] > 0 {
                    result.push(nums[i]);
                    a = i + 1;
                    break;
                }
            }
            for i in b..max {
                if nums[i] < 0 {
                    result.push(nums[i]);
                    b = i + 1;
                    break;
                }
            }
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    struct Arg {
        input: i32,
        expect: String,
    }

    impl Arg {
        fn new(input: i32, expect: String) -> Arg {
            Arg { input, expect }
        }
    }

    #[test]
    fn test_rearrange() {
        let nums = vec![
            28, -41, 22, -8, -37, 46, 35, -9, 18, -6, 19, -26, -37, -10, -9, 15, 14, 31,
        ];
        let result = Solution::rearrange_array(nums);
        eprintln!("rearrange_array={:?}", result);
        let expect = vec![
            28, -41, 22, -8, 46, -37, 35, -9, 18, -6, 19, -26, 15, -37, 14, -10, 31, -9,
        ];
        assert_eq!(expect, result);
    }
}
