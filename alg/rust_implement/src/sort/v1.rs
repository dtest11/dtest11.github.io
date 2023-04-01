pub struct SortV1 {}
impl SortV1 {
    // 冒泡排序：循环n次,每次将最大的数据放到最后一位。
    // 时间复杂度 n *n =n2
    // 空间复杂度: n
    pub fn bubble_sort(nums: &mut Vec<i32>) {
        eprintln!("****bubble sort");
        eprintln!("每一轮循环将，最小或者最大的数据放到第一位:{:?}", nums);
        let l = nums.len() - 1;
        for i in 0..l {
            for j in 0..l - i {
                // eprintln!("max={:?}", j+1);
                if nums[j] > nums[j + 1] {
                    (nums[j], nums[j + 1]) = (nums[j + 1], nums[j]);
                }
            }
            // eprintln!("第{:?}次排序结果{:?}",i,nums);
        }
        // eprintln!("冒泡排序结果:{:#?}", nums);
        eprintln!("冒泡排序结果:{:?}", nums);
    }
    // 合并排序：将数据等分为2半，然后将2半，依次merge排序合并
    // 时间复杂度:O(NlogN)
    // 空间复杂度：O(N)
    // 归并排序算法具有稳定性
    // https://juejin.cn/post/6993926231708139533
    pub fn merge_sort(nums: &Vec<i32>) -> Vec<i32> {
        if nums.len() < 2 {
            return nums.to_vec();
        }
        let mid = nums.len() / 2;
        let left_nums = SortV1::merge_sort(&nums[0..mid].to_vec());
        let right_nums = SortV1::merge_sort(&nums[mid..].to_vec());
        let result = SortV1::merge(left_nums, right_nums);
        result
    }
    pub fn merge(left: Vec<i32>, right: Vec<i32>) -> Vec<i32> {
        let mut i = 0;
        let mut j = 0;
        let max_i = left.len() - 1;
        let max_j = right.len() - 1;
        let mut result: Vec<i32> = Vec::new();

        while i <= max_i && j < max_j {
            if left[i] < right[j] {
                result.push(left[i]);
                i = i + 1;
            } else {
                result.push(right[j]);
                j = j + 1;
            }
        }
        while i <= max_i {
            result.push(left[i]);
            i = i + 1;
        }
        while j <= max_j {
            result.push(right[j]);
            j = j + 1;
        }
        result
    }
    // 快速排序
    // 每次将数据分成2半，将大于pivot的数据移动到右边，小于pivot移动到左边
    // 递归处理
    // 时间复杂度: 平均情况下快速排序的时间复杂度是Θ(nlgn)，最坏情况是n2
    //，但通过随机算法可以避免最坏情况。由于递归调用，快排的空间复杂度是Θ(lgn)
    // 空间复杂度：
    // 是否是稳定排序：
    pub fn quick_sort(nums: &mut Vec<i32>, left: usize, right: usize) {
        if left > right {
            return;
        }
        let pivot = SortV1::get_pivot(nums, left, right);
        SortV1::quick_sort(nums, left, pivot-1);
        SortV1::quick_sort(nums, pivot+1, right);
    }
    pub fn get_pivot(nums: &mut Vec<i32>, mut left: usize, mut right: usize) ->usize{
        let copy_left = left;
        let cur = nums[left];
        while left < right {
            while left < right && nums[right] >= cur {
                right = right - 1;
            }
            while left < right && nums[right] <= cur {
                left = left + 1;
            }
            while left < right {
                (nums[left], nums[right]) = (nums[right], nums[left]);
            }
        }
        (nums[left],nums[copy_left])=(nums[copy_left],nums[left]);
        left
    }
}

#[cfg(test)]
mod tests {
    use super::SortV1;

    struct Args {
        input: Vec<i32>,
        expect: Vec<i32>,
    }
    impl Args {
        fn new(a: Vec<i32>, b: Vec<i32>) -> Self {
            Args {
                input: a,
                expect: b,
            }
        }
    }

    #[test]
    pub fn test_sort() {
        let nums = vec![3, 2, 4, 1];
        let mut args: Vec<Args> = Vec::new();

        args.push(Args::new(vec![3, 2, 4, 1], vec![1, 2, 3, 4]));
        args.push(Args::new(
            vec![0, 1, 7, 3, 4, 9, 8],
            vec![0, 1, 3, 4, 7, 8, 9],
        ));
        args.push(Args::new(
            vec![0, 1, 3, 1, 4, 9, 8],
            vec![0, 1, 1, 3, 4, 8, 9],
        ));
        for a in args {
            let mut input = Vec::from(a.input.clone());
            SortV1::bubble_sort(&mut input);
            assert_eq!(input, a.expect);

            // merge sort
            let input1 = Vec::from(a.input.clone());
            assert_eq!(SortV1::merge_sort(&mut input), a.expect);
        }
    }
}
