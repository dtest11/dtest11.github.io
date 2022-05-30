/***
层级遍历
 */

use std::rc::Rc;
use std::cell::RefCell;
use std::collections::VecDeque;

#[derive(Debug)]
struct TreeNode {
    val: i32,
    left: Option<Rc<RefCell<TreeNode>>>,
    right: Option<Rc<RefCell<TreeNode>>>,
}



#[allow(unused)]
fn level_order(root: Option<Rc<RefCell<TreeNode>>>) ->Vec<Vec<i32>>{
    let mut result = Vec::new();
    if root.is_none() {
        return result;
    }
    let mut queue: VecDeque<Option<Rc<RefCell<TreeNode>>>> = VecDeque::new();

    queue.push_back(root);
    while !queue.is_empty() {
        let mut level_value = vec![];
        let level_size = queue.len();
        for _i in 0..level_size {
            let cur = queue.pop_front();
            if let Some(Some(node)) = cur {
                level_value.push(node.borrow().val);
                if node.borrow().left.is_none() {
                    queue.push_back(node.borrow().left.clone());
                }
                if node.borrow().right.is_none() {
                    queue.push_back(node.borrow().right.clone());
                }
            }
        }
        result.push(level_value);
    }
    result
}