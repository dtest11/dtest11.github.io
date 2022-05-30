use std::collections::{ HashMap};

#[allow(unused)]
fn add(r: &mut HashMap<String, bool>, key: String) -> bool {
    let result = r.insert(key, true);
    return match result {
        None => {
            true
        }
        Some(_i) => {
            false
        }
    };
}


// impl Solution {
//     pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
//         let mut record = HashMap::new();
//         // let mut found = HashSet::new();
//         for r in 0..9 {
//             for c in 0..9 {
//                 let num = board[r][c];
//                 if num != '.' {
//                     let k = format!("{}+{}+r", num, r);
//                     let k2 = format!("{}+{}+c", num, c);
//                     let k3 = format!("{}+{}+{}+b", num, r / 3, c / 3);
//                     // if found.insert(k) || found.insert(k2) || found.insert(k3) {
//                     //     return false;
//                     // }
//                     if add(&mut record, k) || add(&mut record, k2) || add(&mut record, k3) {
//                         return false;
//                     }
//                 }
//             }
//         }
//         true
//     }
// }