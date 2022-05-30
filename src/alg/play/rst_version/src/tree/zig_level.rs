use std::rc::Rc;
use std::cell::RefCell;
use std::collections::VecDeque;


#[derive(Debug)]
struct TreeNode {
    val: i32,
    left: Option<Rc<RefCell<TreeNode>>>,
    right: Option<Rc<RefCell<TreeNode>>>,
}


impl TreeNode {
    #[inline]
    #[allow(unused)]
    fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}
#[allow(unused)]
struct Solution {}


impl Solution {
    #[allow(unused)]
    fn recover_tree(root: Option<Rc<RefCell<TreeNode>>>) {
        if root.is_none() {
            return;
        }
        let mut stack: Vec<Option<Rc<RefCell<TreeNode>>>> = Vec::new();
        let mut copy_cur: Vec<Option<Rc<RefCell<TreeNode>>>> = Vec::new();
        copy_cur.push(root);
        while !copy_cur.is_empty() || !stack.is_empty() { // 一路到地，找到最左边的位置
            while !copy_cur.is_empty() {
                let top = copy_cur.pop();
                if let Some(Some(node)) = top {
                    println!("{}", node.borrow().val);
                    stack.push(Option::from(node.clone()))
                }
            }
            let top = stack.pop();
            if let Some(Some(node)) = top {
                println!("{}", node.borrow().val);
                copy_cur.push(node.borrow().right.clone());
            }
        }
    }

    #[allow(unused)]
    pub fn zigzag_level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        if root.is_none() {
            return Vec::new();
        }

        let mut result = Vec::new();
        let mut queue = VecDeque::new();
        queue.push_back(root.clone());
        let mut left_to_right = false;
        while !queue.is_empty() {
            let size = queue.len();
            let mut tmp = Vec::new();
            for _i in 0..size {
                let top = queue.pop_front();
                if let Some(Some(node)) = top {
                    tmp.push(node.borrow().val);

                    if !node.borrow().left.is_none() {
                        queue.push_back(node.borrow().left.clone());
                    }
                    if !node.borrow().right.is_none() {
                        queue.push_back(node.borrow().right.clone());
                    }
                }
            }
            if left_to_right {
                tmp.reverse();
            }
            left_to_right = !left_to_right;
            result.push(tmp);
        }
        result
    }
}
