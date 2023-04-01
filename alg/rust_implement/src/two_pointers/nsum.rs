pub struct Solution;

impl Solution {
    // four_sum：四数之和，一样的双指针，每一层循环需要注意删除重复的数据
    pub fn four_sum(v: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut nums = v.to_vec();
        let mut res = Vec::new();

        nums.sort();
        let max = nums.len();
        for a in 0..max {
            if a > 0 && nums[a] == nums[a - 1] {
                continue;
            }
            for b in a + 1..max {
                if b > a + 1 && nums[b] == nums[b - 1] {
                    continue;
                }
                let mut l = b + 1;
                let mut r = max - 1;
                while l < r {
                    let sum = nums[l] as i64 + nums[r] as i64 + nums[a] as i64 + nums[b] as i64;
                    if sum < target as i64 {
                        l = l + 1;
                    } else if sum > target as i64 {
                        r = r - 1;
                    } else {
                        // sum == target
                        res.push(vec![nums[a], nums[b], nums[l], nums[r]]);
                        l = l + 1;
                        r = r - 1;
                        while l < r && nums[l - 1] == nums[l] {
                            l = l + 1;
                        }
                        while l < r && nums[r] == nums[r + 1] {
                            r = r - 1;
                        }
                    }
                }
            }
        }

        res
    }
    // https://leetcode.cn/problems/3sum/solution/san-shu-zhi-he-by-leetcode-solution/
    // 只有旁边的位置相等，才进行查询，比如数据是[1,3,3,-1,-2] 如果不把3的重复跳过，会出现重复[3,-1,2]
    pub fn three_sum(v: Vec<i32>, target: i32, l: usize) -> Vec<Vec<i32>> {
        let mut nums = v.to_vec();
        nums.sort();

        let mut res = Vec::new();
        let len = nums.len() - 1;
        for i in l..len {
            if i != l && nums[i - 1] == nums[i] {
                continue;
            }

            let temp = Solution::two_sum(nums.to_vec(), target - nums[i], i + 1);
            for t in temp {
                let mut t1 = t;
                t1.push(nums[i]);
                res.push(t1);
            }
        }

        res
    }

    // two_sum_with_target 2数之和, 双指针，low=0,high=nums.len()-1, 开始移动
    // sum > target, 移动high
    // sum < target 一定low
    // sum == target ,同时消除重复的数据
    pub fn two_sum(nums: Vec<i32>, target: i32, mut l: usize) -> Vec<Vec<i32>> {
        let mut result = Vec::new();
        let mut r = nums.len() - 1;
        while l < r {
            let sum = nums[l] + nums[r];
            if sum < target {
                l = l + 1;
            } else if sum > target {
                r = r - 1;
            } else {
                // sum == target
                result.push(vec![nums[l], nums[r]]);
                l = l + 1;
                r = r - 1;
                // 删除重复的数据
                while l < r && nums[l - 1] == nums[l] {
                    l = l + 1;
                }
                // 删除重复的数据
                while l < r && nums[r] == nums[r + 1] {
                    r = r - 1;
                }
            }
        }
        result
    }

    // 最靠近target的数据：即为差值的绝对值最小
    pub fn three_sum_closest(v: Vec<i32>, target: i32) -> i32 {
        let mut prev_sum: i32 = v[0] + v[1] + v[2];
        let mut nums = v.to_vec();
       nums.sort();
        for i in 0..nums.len() - 1 {
            if i != 0 && nums[i - 1] == nums[i] {
                continue;
            }
            let mut l = i + 1;
            let mut r = nums.len() - 1;
            while l < r {
                let sum = nums[l] + nums[r] + nums[i];
                if sum > target {
                    r = r - 1;
                } else if sum < target {
                    l = l + 1;
                } else {
                    return sum;
                }
                if (sum - target).abs() < (prev_sum - target).abs() {
                    prev_sum = sum;
                }
            }
        }
        prev_sum
    }
}

pub struct NsumV2;

impl NsumV2 {
    pub fn four_sum(v: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        return NsumV2::n_sum(v, target, 4);
    }
    pub fn n_sum(v: Vec<i32>, target: i32, n: i32) -> Vec<Vec<i32>> {
        let mut nums = v.to_vec();
        nums.sort();
        return NsumV2::n_sum_util(nums.to_vec(), target as i64, 0, n);
    }

    pub fn n_sum_util(nums: Vec<i32>, target: i64, mut l: usize, n: i32) -> Vec<Vec<i32>> {
        let mut result = Vec::new();
        let max = nums.len();
        if n == 2 {
            // two sum
            let mut r = nums.len() - 1;
            while l < r {
                let sum = nums[l] as i64 + nums[r] as i64;
                if sum < target {
                    l = l + 1;
                } else if sum > target {
                    r = r - 1;
                } else {
                    // sum == target
                    result.push(vec![nums[l], nums[r]]);
                    l = l + 1;
                    r = r - 1;
                    while l < r && nums[l - 1] == nums[l] {
                        l = l + 1;
                    }
                    while l < r && nums[r] == nums[r + 1] {
                        r = r - 1;
                    }
                }
            }
            return result;
        }
        for i in l..nums.len() {
            if i > l && nums[i - 1] == nums[i] {
                continue;
            }
            let temp = NsumV2::n_sum_util(nums.to_vec(), target - nums[i] as i64, i + 1, n - 1);
            for t in temp {
                let mut t1 = t;
                t1.push(nums[i]);
                result.push(t1)
            }
        }

        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    struct Args {
        nums: Vec<i32>,
        expect: Vec<Vec<i32>>,
        target: i32,
        n: i32,
    }

    #[test]
    fn test_three_sum() {
        let nums = vec![2, 2, 2, 2, 2];
        let arg = Args {
            expect: vec![vec![2, 2, 2]],
            nums: vec![2, 2, 2, 2, 2],
            target: 6,
            n: 0,
        };
        let result = Solution::three_sum(arg.nums, arg.target, 0);
        assert_eq!(result, arg.expect);
    }

    #[test]
    fn test_n_sum() {
        let empty: Vec<Vec<i32>> = Vec::new();
        let args = vec![
            Args {
                nums: vec![2, 2, 2, 2, 2],
                expect: vec![vec![2, 2, 2, 2]],
                target: 8,
                n: 4,
            },
            Args {
                nums: vec![1000000000, 1000000000, 1000000000, 1000000000],
                expect: empty,
                target: -294967296,
                n: 4,
            },
            Args {
                target: 0,
                expect: vec![vec![1, 2, -1, -2], vec![1, 1, -1, -1]],
                nums: vec![-2, -1, -1, 1, 1, 2, 2],
                n: 4,
            },
        ];
        for arg in args {
            let result = NsumV2::n_sum(arg.nums, arg.target, arg.n);
            assert_eq!(result, arg.expect);
        }
    }

    #[test]
    fn test_four_sum() {
        let empty: Vec<Vec<i32>> = Vec::new();

        let args: Vec<Args> = vec![
            Args {
                nums: vec![2, 2, 2, 2, 2],
                expect: vec![vec![2, 2, 2, 2]],
                target: 8,
                n: 0,
            },
            Args {
                nums: vec![-2, -1, -1, 1, 1, 2, 2],
                expect: vec![vec![-2, -1, 1, 2], vec![-1, -1, 1, 1]],
                target: 0,
                n: 0,
            },
            Args {
                nums: vec![1000000000, 1000000000, 1000000000, 1000000000],
                expect: empty,
                n: 0,
                target: -294967296,
            },
        ];
        let empty: Vec<Vec<i32>> = Vec::new();
        for arg in args {
            let result = Solution::four_sum(arg.nums, arg.target);
            assert_eq!(result, arg.expect);
        }
    }
}
