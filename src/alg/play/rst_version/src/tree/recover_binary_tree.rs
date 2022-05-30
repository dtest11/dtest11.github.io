
#[allow(unused)]
struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug)]
struct Node {
    val: i32,
    left: Option<Rc<RefCell<Node>>>,
    right: Option<Rc<RefCell<Node>>>,
}


impl Node {
    #[inline]
    #[allow(unused)]
    fn new(val: i32) -> Self {
        Node {
            val,
            left: None,
            right: None,
        }
    }
}

#[allow(unused)]
impl Solution {
    pub fn recover_tree(_root: Option<Rc<RefCell<Node>>>) {
        // 中序遍历
    }
}