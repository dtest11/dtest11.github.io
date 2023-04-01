#![allow(dead_code)]
#![allow(unused_variables)]

use std::cmp::Ordering;

// 冒泡排序，一趟下来把合适的元素移动到最后一个
// n*n的时间复杂度
pub fn bubble_sort(data: &mut [i32]) {
    println!("冒泡排序");
    for i in 0..data.len() - 1 {
        let mut j = 1;
        while j <= data.len() - 1 - i {
            if data[j - 1] > data[j] {
                data.swap(j - 1, j);
            }
            j += 1;
        }
    }
}

// merge sort:
// 1. merge
// 2. sort
// 先将数据分成2部分：递归，
// 然后将数据merge
pub fn merge_sort(array: Vec<i32>) -> Vec<i32> {
    println!("归并排序");
    if array.len() < 2 {
        array
    } else {
        let mid = array.len() / 2;
        let left = merge_sort(array[0..mid].to_vec());
        let right = merge_sort(array[mid..].to_vec());
        let result = merge(left, right);
        result
    }
}

fn merge(left: Vec<i32>, right: Vec<i32>) -> Vec<i32> {
    let mut result = Vec::new();
    let mut i = 0;
    let mut j = 0;
    while i < left.len() && j < right.len() {
        match left[i].cmp(&right[j]) {
            Ordering::Less => {
                result.push(right[j]);
                j += 1;
            }
            Ordering::Greater => {
                result.push(left[i]);
                i += 1;
            }
            Ordering::Equal => {
                result.push(left[i]);
                i += 1;
                result.push(right[j]);
                j += 1;
            }
        }
    }
    while i < left.len() {
        result.push(left[i]);
        i += 1;
    }
    while j < right.len() {
        result.push(right[j]);
        j += 1;
    }
    result
}

// 思路： 找到一个临界点，将所有大于该临界点的数据移动到左边
// 将所有大于临界点的值移动到right
fn quick_sort(data: &mut Vec<i32>) {
    quick_sort_util(data, 0, data.len());
}

fn quick_sort_util(data: &mut Vec<i32>, left: usize, right: usize) {
    if left < right {
        let pivot = pivot_helper(data, left, right);
        quick_sort_util(data, left, pivot - 1);
        quick_sort_util(data, pivot + 1, right);
    }
}

fn pivot_helper(data: &mut Vec<i32>, mut left: usize, mut right: usize) -> usize {
    let point = data[left];
    let copy_left = left;
    while left < right {
        while left < right && data[right] > point {
            right -= 1;
        }
        while left < right && data[left] < point {
            left += 1;
        }
        if left < right {
            data.swap(left, right);
        }
    }
    data.swap(left, copy_left);
    left
}

#[cfg(test)]
mod tests {
    use crate::sort::sort::bubble_sort;
    use crate::sort::sort::merge_sort;

    #[test]
    fn it_works() {
        let mut data = [1, 5, 2, 3, 1, 8];
        bubble_sort(&mut data);
        println!("{:?}", data);
    }
    #[test]
    fn merge() {
        // let  data = [1, 5, 2, 3, 1, 8];
        let data = [3, 2, 4, 1];
        let result = merge_sort(data.to_vec());
        println!("merge output:{:?}", result);
    }
}
