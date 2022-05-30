use std::cell::RefCell;
use std::rc::Rc;

pub struct TreeNode {
    val: i32,
    left: Option<Rc<RefCell<TreeNode>>>,
    right: Option<Rc<RefCell<TreeNode>>>,
}

struct Solution {}

impl Solution {
    #[allow(unused)]
    pub fn is_symmetric_recurse(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if root.is_none() {
            return true;
        }
        let n = match root {
            Some(node) => node,
            None => return false,
        };
        return Solution::helper(n.borrow().left.clone(), n.borrow().right.clone());
    }

    #[allow(unused)]
    fn helper(left: Option<Rc<RefCell<TreeNode>>>, right: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if left.is_none() && right.is_none() {
            return true;
        }
        if left.is_none() || right.is_none() {
            return false;
        }
        let a = match left {
            Some(a_node) => a_node,
            None => return false,
        };
        let b = match right {
            Some(a_node) => a_node,
            None => return false,
        };
        if a.borrow().val != b.borrow().val {
            return false;
        }
        return Solution::helper(a.borrow().left.clone(), b.borrow().right.clone())
            || Solution::helper(a.borrow().right.clone(), b.borrow().left.clone());
    }

    #[allow(unused)]
    pub fn is_symmetric(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if root.is_none() {
            return true;
        }
        let n = match root {
            Some(node) => node,
            None => return false,
        };
        return Solution::hello(n.borrow().left.clone(), n.borrow().right.clone());
    }

    #[allow(unused)]
    fn hello(left: Option<Rc<RefCell<TreeNode>>>, right: Option<Rc<RefCell<TreeNode>>>) -> bool {
        let mut stack: Vec<Option<Rc<RefCell<TreeNode>>>> = Vec::new();
        stack.push(left);
        stack.push(right);

        while !stack.is_empty() {
            let p = stack.pop();
            let q = stack.pop();


            if p.is_none() && q.is_none() {
                continue;
            }

            let a = match p {
                Some(a_node) => a_node,
                None => return false,
            };


            let b = match q {
                Some(a_node) => a_node,
                None => return false,
            };

            if a.is_none() && b.is_none() {
                continue;
            }

            let f = match a {
                Some(a_node) => a_node,
                None => return false
            };
            let g = match b {
                Some(a_node) => a_node,
                None => return false
            };


            if f.borrow().val != g.borrow().val {
                return false;
            }
            let gg = g.borrow().left.clone();
            stack.push(f.borrow().left.clone());
            stack.push(g.borrow().right.clone());

            stack.push(f.borrow().right.clone());
            stack.push(gg)
        }
        true
    }
}
