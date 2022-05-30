#[allow(unused)]
pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
    let mut position: usize = 0;

    if nums.is_empty() {
        return position as i32;
    }
    for i in 1..nums.len() {
        // let current: i32 = nums[i];
        // let prev = nums[i - 1];
        if nums[i] != nums[i - 1] {
            position = position + 1;
            nums[position] = nums[i];
        }
    }
    let result = position + 1;
    result as i32
}



