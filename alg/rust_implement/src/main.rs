#![allow(dead_code)]
#![allow(unused_variables)]


mod diget;
mod sort;
mod two_pointers;

use crate::sort::v1;

fn main() {
    let data =vec![1,2];
    eprintln!("{:?}",&data[1..]);
    let mut nums = vec![0, 1, 7, 3, 4, 9, 8];
    v1::SortV1::bubble_sort(&mut nums);
    // Sort::merge_sort(&nums);
    // eprintln!("{:?}", nums);
    // let res = Solution::reverse_words(String::from("Hello,  world!"));
}



