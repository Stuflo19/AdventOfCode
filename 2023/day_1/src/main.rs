use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::HashMap;
fn main() {
    // part_1("./input.txt");
    part_2("./input.txt");
}

fn part_1(filename: &str) {
    let mut total = 0;

    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(input) = line {
                let mut left = 'x';
                let mut right = 'x';
                for char in input.chars() {
                    if char.is_digit(10) && left == 'x' {
                        left = char;
                        right = char;
                    } else if char.is_digit(10) {
                        right = char;
                    }
                }

                let result = left.to_string() + "" + &*right.to_string();
                total += result.parse::<i32>().unwrap();
            }
        }
    }

    println!("{}", total);
}

fn part_2(filename: &str) {
    let mut total = 0;
    let mut words: HashMap<String, char> = HashMap::new();

    words.insert(String::from("one"), '1');
    words.insert(String::from("two"), '2');
    words.insert(String::from("three"), '3');
    words.insert(String::from("four"), '4');
    words.insert(String::from("five"), '5');
    words.insert(String::from("six"), '6');
    words.insert(String::from("seven"), '7');
    words.insert(String::from("eight"), '8');
    words.insert(String::from("nine"), '9');


    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(input) = line {
                let mut left = 'x';
                let mut right = 'x';
                let mut strBuilder :String = String::from("");

                for c in input.chars() {
                    strBuilder = strBuilder + c.to_string().as_str();

                    if c.is_digit(10) {
                        left = c;
                        break;
                    }

                    let mut set = false;
                    for k in words.keys() {
                        // println!("key: {}", k);
                        if strBuilder.contains(k) {
                            left = *words.get(k).unwrap();
                            // println!("Left: {}", left);
                            set = true;
                            break;
                        }
                    }
                    if set { break; }
                }

                strBuilder = String::from("");

                for c in input.chars().rev() {
                    strBuilder = c.to_string() + &*strBuilder;
                    println!("str: {}", strBuilder);
                    if(c.is_digit(10)) {
                        right = c;
                        break;
                    }

                    let mut set = false;
                    for k in words.keys() {
                        if strBuilder.contains(k) {
                            right = *words.get(k).unwrap();
                            set = true;
                            break;
                        }
                    }
                    if set { break; }
                }

                let result = left.to_string() + "" + &*right.to_string();
                println!("Result: {}", result);
                total += result.parse::<i32>().unwrap();
            }
        }
    }

    println!("Total: {}", total);
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
