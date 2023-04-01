#![allow(dead_code)]
#![allow(unused_variables)]

pub struct Node {
    pub val: i32,
    pub next: Option<Box<Node>>,
}
pub struct LinkList {
    head: Option<Box<Node>>,
}
/**
 * Box: 会将数据存储到堆上，可以定义循环递归引用的数据，只需要少量的指针就可以做关联
 * 使用 Rc<T> 允许一个值有多个所有者，引用计数则确保只要任何所有者依然存在其值也保持有效。
 Rc<T> 允许相同数据有多个所有者；Box<T> 和 RefCell<T> 有单一所有者。
Box<T> 允许在编译时执行不可变或可变借用检查；Rc<T>仅允许在编译时执行不可变借用检查；RefCell<T> 允许在运行时执行不可变或可变借用检查。
因为 RefCell<T> 允许在运行时执行可变借用检查，所以我们可以在即便 RefCell<T> 自身是不可变的情况下修改其内部的值。
https://rust-unofficial.github.io/too-many-lists/second.html
 */
impl LinkList {
    pub fn new(val: i32) -> Self {
        LinkList { head: None }
    }
    pub fn push(&mut self, val: i32) {
        let new_node = Box::new(Node {
             val,
            next: self.head.take(),
        });
        self.head = Some(new_node);
    }
    //
    // pub fn pop(&mut self) -> Option<i32> {
    //     match self.head.take() {
    //         Node => None,
    //     }
    // }
}


