use std::fs;
use std::num::ParseIntError;

fn get_num(line: String) -> Result<i32, ParseIntError> {
    let num: String = line.chars().filter(|c| c.is_numeric()).collect();

    let first = num.chars().next().unwrap();
    let last = num.chars().last().unwrap();
    format!("{}{}", first, last).parse()
}

fn main() {
    let file = fs::read_to_string("day01.txt").expect("No such file");

    let mut sum: i32 = 0;
    for line in file.lines() {
        let num = match get_num(line.to_string()) {
            Ok(n) => n,
            Err(err) => {
                println!("Error parsing number: {}", err);
                continue;
            }
        };

        sum += num;
    }

    println!("{}", sum);
}
