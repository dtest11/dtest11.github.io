use std::collections::HashMap;
use std::env;
use std::fs::File;
use std::fs::OpenOptions;
use std::io::prelude::*;
use std::io::{BufReader};
use std::io::{ Write};
use std::process::Command;
use std::io;

static JUMP_FILE: &str = "/home/code/Desktop/happy_coding/jj/target.txt";

fn main() {
    let args: Vec<String> = env::args().collect();
    // println!("{:?}", args);

    if args.len() == 3 {
        let input = args.last().unwrap().to_string();
        // println!("输入 {:}", input);

        let prev_data = load_data(input);
        let mut sorted: Vec<(&String, &i32)> = prev_data.iter().collect();
        sorted.sort_by(|a, b| b.1.cmp(a.1));

        // println!("排序后的 {:?}", sorted);
        write_data(sorted);
    }
    if args.len() == 2 {
        match find_dir(args.last().unwrap()) {
            None => (),
            Some(d) => {
                // cd(d);
                io::stdout().write_all(d.as_bytes());
            }
        };
    }
}

fn cd(target: String) {
    println!("跳转目录:{}", target);
    Command::new("cd")
        .arg(&target)
        .spawn()
        .expect("sh command failed to start");
}

fn find_dir(target: &String) -> Option<String> {
    let file = match File::open(JUMP_FILE) {
        Err(why) => panic!("not found:{:}", why),
        Ok(file) => file,
    };

    let reader = BufReader::new(file);
    for l in reader.lines() {
        let copy_line = l.unwrap();
        let array: Vec<_> = copy_line.split(":").collect();
        let dir = array.first().unwrap().to_string();
        let up_dir = dir.to_uppercase();
        let up_target = target.to_uppercase();
        if up_dir.contains(&up_target) || up_dir.eq(&up_target) {
            return Some(dir);
        }
    }
    None
}

fn load_data(input: String) -> HashMap<String, i32> {
    let mut record: HashMap<String, i32> = HashMap::new();
    // println!("创建文件{}",JUMP_FILE);
    let new_file = OpenOptions::new().create(true).read(true).write(true).open(JUMP_FILE);
    let file = match new_file{
        Err(why) => {
            panic!("打开文件失败:{}:{}", JUMP_FILE,why);
        //    return  record;
        }
        Ok(file) => file,
    };

    let reader = BufReader::new(file);
    for line in reader.lines() {
        match line{
            Ok(content) => {
                let copy_line: String =content;
                let l: Vec<_> = copy_line.split(":").collect();
                record.insert(
                    l.first().unwrap().parse().unwrap(),
                    l.last().unwrap().parse().unwrap(),
                );
            }
            Err(_) => {
                break
            }
        }

    }
    // println!("new record:{}",input);

    if let Some(x) = record.get_mut(&input) {
        *x = *x + 1;
    } else {
        record.insert(input, 1);
    }

    record
}

fn write_data(data: Vec<(&String, &i32)>) {
    let mut file = match OpenOptions::new()
        .write(true)
        .truncate(true)
        .create(true)
        .open(JUMP_FILE)
    {
        Err(why) => panic!("无法创建文件{},{}", why, JUMP_FILE),
        Ok(file) => file,
    };
    for (k, v) in data.iter() {
        // println!("写入新的记录{}:{},{}", JUMP_FILE, k, v);
        match writeln!(file, "{}:{}", k, v) {
            Err(why) => eprintln!("写入失败:,{},{}", why, JUMP_FILE),
            Ok(_) => println!("ok"),
        }
    }
}
