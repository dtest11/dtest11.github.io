fn main() {
    println!("Hello, world!");
    let mut data: Vec<i32> = vec![10, 2, 1, 5, 3];
    bubble(&mut data);
    println!("{:?}", data);
}

fn bubble(data: &mut Vec<i32>) {
    for i in 0..data.len() {
        for j in 1..data.len() - i {
            if data[j - 1] > data[j] {
                let tmp = data[j - 1];
                data[j - 1] = data[j];
                data[j] = tmp;
            }
        }
    }
}
