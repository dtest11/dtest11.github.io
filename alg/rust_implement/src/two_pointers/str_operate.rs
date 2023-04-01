#![allow(dead_code)]
#![allow(unused_variables)]
#![warn(unused_assignments)]

pub struct Solution {}

impl Solution {
    pub fn flip_and_invert_image(image: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut new_image: Vec<Vec<i32>> = vec![vec![0; image.len()]; image.len()];
        for i in 0..image.len() {
            for j in 0..image.len() {
                new_image[i][j] = image[i][j];
            }
        }
        // flipping
        for i in 0..image.len() {
            let mut left = 0;
            let mut right = image.len() - 1;
            while left < right {
                let temp = new_image[i][right];
                new_image[i][right] = new_image[i][left];
                new_image[i][left] = temp;
                left += 1;
                right -= 1;
            }
        }
        // invert
        for i in 0..image.len() {
            for j in 0..image.len() {
                if new_image[i][j] == 0 {
                    new_image[i][j] = 1
                } else {
                    new_image[i][j] = 0
                }
            }
        }
        new_image
    }
    /**
         * i < j < k,
    nums[j] - nums[i] == diff, and
    nums[k] - nums[j] == diff.
    Return the number of unique arith
     */
    pub fn arithmetic_triplets(nums: Vec<i32>, diff: i32) -> i32 {
        let mut result = 0;
        let len = nums.len();
        for i in 0..len {
            for j in i + 1..len {
                for k in i + 2..len {
                    if nums[j] - nums[i] == diff && nums[k] - nums[j] == diff {
                        result += 1;
                        // println!("{},{},{}", i, j, k);
                    }
                }
            }
        }
        result
    }

    pub fn reverse_str(s: String) -> String {
        let mut letters: Vec<char> = s.chars().collect();
        let l = letters.len();
        let mut left = 0;
        let mut right = l - 1;
        while left < right {
            let temp = letters[right];
            letters[right] = letters[left];
            letters[left] = temp;
            left += 1;
            right -= 1;
        }
        let mut buf = String::new();
        for i in letters.iter() {
            buf.push(*i);
        }
        buf
    }

    pub fn reverse_words(s: String) -> String {
        let letters = s.chars().collect::<Vec<char>>();
        let mut result = "".to_string();
        // let mut slow = 0;
        let mut fast = 0;
        let mut cache_char = Vec::new();
        while fast < letters.len() {
            let cur = letters[fast];
            if cur == ' ' {
                // 遇到空格，前面的数据 是一组合法的字符串，需要翻转
                let mut reverse_str: String = cache_char.iter().collect();
                reverse_str = Solution::reverse_str(reverse_str);

                result.push_str(&reverse_str);
                cache_char.clear(); // 删除历史数据
                                    // 继续遍历字符串，直到下一个字符不是空格，将slow的位置移动到此处
                while fast < letters.len() && letters[fast] == ' ' {
                    result.push(letters[fast]);
                    fast += 1;
                }
                cache_char.push(letters[fast]);
                // slow = fast;
            } else {
                // 字符串不为空
                cache_char.push(cur);
            }
            fast = fast + 1
        }
        if cache_char.len() != 0 {
            let mut reverse_str: String = cache_char.iter().collect();
            reverse_str = Solution::reverse_str(reverse_str);
            result.push_str(&reverse_str);
        }
        return result;
    }
}

// 3个指针， 第一个i,第二个从i+1开始，第三个从i+2开始遍历
// cargo test -- --color always --nocapture
#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_bad_add() {
        let nums = vec![0, 1, 4, 6, 7, 10];
        let diff = 3;
        let result = Solution::arithmetic_triplets(nums, diff);
        assert_eq!(result, 2);
    }

    #[test]
    fn test_reverse() {
        let str = String::from("hello");
        let result = Solution::reverse_str(str);
        assert_eq!(result, String::from("olleh"));

        let str = String::from("Let's take LeetCode contest");
        let expect = String::from("s'teL ekat edoCteeL tsetnoc");
        assert_eq!(Solution::reverse_words(str), expect)
    }

    #[test]
    fn test_flipp_image() {
        let arrays = vec![vec![0; 10]; 10];
        let res = Solution::flip_and_invert_image(arrays.clone());
        for i in 0..arrays.len() {
            for j in 0..arrays.len() {
                // println!("{:?}", arrays[i][j])
            }
        }
    }
}
