
#[allow(unused)]
pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
    if nums.is_empty() {
        return 0;
    }
    if nums.len() == 1 {
        if nums[0] == val {
            return 0;
        }
        return 1;
    }
    let mut i: usize = 0;
    for index in 0..nums.len() {
        if nums[index] != val {
            nums[i] = nums[index];
            i = i + 1;
        }
    }
    let result: i32 = i as i32;
    result
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_remove_elements() {
        let mut input: Vec<i32> = vec![1, 1, 1, 2, 2, 3];
        let res = remove_element(&mut input, 1);
        println!("{}", res);
    }
}
