use std::rc::Rc;
use std::cell::RefCell;

#[allow(unused)]
 struct TreeNode {
    val: i32,
    left: Option<Rc<RefCell<TreeNode>>>,
    right: Option<Rc<RefCell<TreeNode>>>,
}
//
// // pre_order_recurse 前序遍历 root-left-right
// fn pre_order_recurse(root: Option<Rc<RefCell<TreeNode>>>) {
//     if let Some(r) = root {
//         println!("val:{}", r.borrow().take().val);
//         pre_order_recurse(r.borrow().take().left);
//         pre_order_recurse(r.borrow().take().right);
//     }
// }