mod graph;
mod sliding_window;
mod sort;
mod dfs;
use crate::sliding_window::window;

fn main() {
    let result = window::length_of_longest_substring("hello".to_string());
    println!("{}", result);
    let mut t = Vec::new();
    t.push(50);
    t.push(30);
    t.push(20);
    let mut h = Vec::new();
    h.push(t);
    for i in &h[0] {
        println!("{:?}", i);
    }
}
