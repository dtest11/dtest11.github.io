use std::collections::HashMap;
use std::collections::LinkedList;

#[allow(unused)]
pub fn next_greater_element(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    let mut result: Vec<i32> = Vec::new();
    let mut record: HashMap<i32, i32> = HashMap::new();
    let mut r_stack: LinkedList<i32> = LinkedList::new();
    for i in 0..nums2.len() {
        let cur = nums2[i];
        loop {
            if r_stack.len() > 0 && r_stack.front().unwrap() < &cur {
                record.insert(r_stack.pop_front().unwrap(), cur);
            } else {
                break;
            }
        }
        r_stack.push_front(cur);
    }

    for e in 0..nums1.len() {
        let found = record.get(&nums1[e]);
        match found {
            None => {
                result.push(-1);
            }
            Some(val) => {
                result.push(*val)
            }
        }
    }
    result
}