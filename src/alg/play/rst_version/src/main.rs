mod strings;
mod array;
mod tree;

fn main() {
    println!("Hello, world!");
    let a ='a';

    for i in 0..9{
        
        println!("{}-{}",i,a);
    }
    let tweet =Article{
        headline: "hello".parse().unwrap(),
    };
    tweet.summarize();
}


pub trait Summary {
    fn summarize(&self) -> String;
}

pub struct Article {
    pub headline:String,
}


impl Summary for Article {
    fn summarize(&self) -> String {
        println!("{}","hello summarize");
        format!("{}","hello")
    }
}