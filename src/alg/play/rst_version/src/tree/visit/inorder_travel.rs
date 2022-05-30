use std::rc::Rc;
use std::cell::RefCell;


#[derive(Debug)]
struct TreeNode {
    val: i32,
    left: Option<Rc<RefCell<TreeNode>>>,
    right: Option<Rc<RefCell<TreeNode>>>,
}

/**
自动
**/
#[allow(unused)]
fn inorder_travel(root: Option<Rc<RefCell<TreeNode>>>) {
    println!("中序遍历: left-root-right");
    if root.is_none() {
        return;
    }
    let mut stack = Vec::new();
    let mut node = root;
    while !node.is_none() || !stack.is_empty() {
        while !node.is_none() {
            stack.push(node.clone());
            if let Some(n) = node.clone() {
                node = n.borrow().left.clone();
            }
        }
        // 到达最left, 然后从底部向上移动
        let top = stack.pop();
        if let Some(Some(ttt)) = top {
            println!("{}", ttt.borrow().val);
            node = ttt.borrow().right.clone();
        }
    }
}

/**

root -left -right
从root节点一路沿着left 开始走
cur = root;
stack
while cur !=nil || stack !=nil
**/
#[allow(unused)]
fn preorder_travel(root: Option<Rc<RefCell<TreeNode>>>){
    println!("preorder序遍历: root -left -right");
    if root.is_none() {
        return;
    }
    let mut stack = Vec::new();
    let mut node = root;
    while !node.is_none() || !stack.is_empty() {
        while !node.is_none() {
            stack.push(node.clone());
            if let Some(n) = node.clone() {
                node = n.borrow().left.clone();
                println!("{}", n.borrow().val);

            }
        }
        // 到达最left, 然后从底部向上移动
        let top = stack.pop();
        if let Some(Some(ttt)) = top {

            node = ttt.borrow().right.clone();
        }
    }
}
#[cfg(test)]
mod tests {
    #[allow(unused)]
    use crate::tree::visit::inorder_travel::{TreeNode, inorder_travel,preorder_travel};
    use std::rc::Rc;
    use std::cell::RefCell;

    #[test]
    fn hello() {
        println!("hello world");
        let left = TreeNode {
            val: 1,
            left: None,
            right: None,
        };
        let right = TreeNode {
            val: 3,
            left: None,
            right: None };
        let root = TreeNode {
            val: 2,
            left: Option::from(Rc::new(RefCell::new(left))),
            right: Option::from(Rc::new(RefCell::new(right))),
        };
        let v = Option::from(Rc::new(RefCell::new(root)));
        println!("expect 1,2,3");
        inorder_travel(v);
        println!("expect 2,1,3");
        // preorder_travel(v);
    }
}



