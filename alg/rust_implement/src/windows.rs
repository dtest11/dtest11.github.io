use std::collections::HashMap;

pub fn length_of_longest_substring(s: String) -> i32 {
    let mut left = 0;
    let mut right = 0;
    let mut map: HashMap<u8, i32> = HashMap::new();
    let mut result = 0;
    let chars = s.into_bytes();
    while right <= chars.len() - 1 {
        let k = chars[right];
        map.entry(chars[right]).and_modify(|e| { *e += 1 }).or_insert(1);
        right += 1;
        while *map.get(&k).unwrap() > 1 {
            let k = chars[left];
            map.entry(k).and_modify(|e| { *e -= 1 });
            left += 1;
        }
        let cur = right - left;
        if cur > result {
            result = cur;
        }
    }

    result as i32
}





#[cfg(test)]
mod tests {
    struct Input {
        name: String,
        args: String,
        expect: i32,
    }

    impl Input {
        fn new(name: String, args: String, expect: i32) -> Input {
            Input { name, args, expect }
        }
    }

    #[test]
    fn it_works() {
        let a = Input::new("1".to_string(), "abcabcbb".to_string(), 3);
        let b = Input::new("2".to_string(), "bbbbb".to_string(), 1);
        let c = Input::new("3".to_string(), "pwwkew".to_string(), 3);
        let  array: [Input; 3] = [a,b,c];
        array.map(|e|{
            let result = crate::window::length_of_longest_substring(e.args);
            assert_eq!(e.expect, result);
            println!("{}", e.name);
        });
    }
}