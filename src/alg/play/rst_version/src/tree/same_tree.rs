use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug)]
pub struct TreeNode {
    val: i32,
    left: Option<Rc<RefCell<TreeNode>>>,
    right: Option<Rc<RefCell<TreeNode>>>,
}

struct Solution{}

impl Solution {
    #[allow(unused)]
    fn is_same_tree(p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>) -> bool {
        // base case
        if p.is_none() && q.is_none() {
            return true;
        }
        // p== nil q !=nil
        if p.is_none() && !q.is_none() {
            return false;
        }
        // p ==nil && q !=nil
        if !p.is_none() && q.is_none() {
            return false;
        }

        let a = match p {
            Some(pnode) => pnode,
            None => return false,
        };

        let b = match q {
            Some(qnode) => qnode,
            None => return false,
        };
        
        let p_val = a.borrow().val;
        let q_val = b.borrow().val;
        if p_val != q_val {
            return false;
        }

        let left_p = a.borrow().left.clone();
        let left_q = b.borrow().left.clone();

        let right_p = a.borrow().right.clone();
        let right_q = b.borrow().right.clone();
        return Solution::is_same_tree(left_p, left_q) && Solution::is_same_tree(right_p, right_q);
    }
}
