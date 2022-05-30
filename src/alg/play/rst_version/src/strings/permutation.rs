#[allow(unused)]
pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut result: Vec<Vec<i32>> = vec![];
    let length = nums.len();
    if nums.is_empty() {
        result
    } else if length == 1 {
        result.push(nums);
        result
    } else {
        permute_helper(&mut result, &mut nums.clone(), 0, length);
        result
    }
}

#[allow(unused)]
pub fn permute_helper(result: &mut Vec<Vec<i32>>, nums: &mut Vec<i32>, left: usize, length: usize) {
    if left == length {
        result.push(nums.clone());
        return;
    }
    for i in left..length {
        if i > left && nums[i] == nums[i - 1] {
            continue; // skip duplicate
        }
        // swap element 
        nums.swap(left, i);
        permute_helper(result, nums, left + 1, length);
        nums.swap(left, i);
    }
}

#[cfg(test)]
mod tests {
    use crate::strings::permutation::permute;

    #[test]
    fn permutation_check() {
        let input: Vec<i32> = vec![1, 2, 3, 4];
        let res = permute(input);
        println!("{:?},{}", res, res.len());
    }
}