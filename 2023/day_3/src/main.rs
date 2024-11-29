use std::collections::HashSet;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::path::Path;

fn main() {
    // part_1("input.txt");
    part_2("input.txt");
}

fn part_1(filename: &str) {
    let mut total = 0;

    let lines = read_lines(filename);
    for (idx, line) in lines.iter().enumerate() {
        let mut temp: String = String::from("");
        let mut starting_symbol = false;
        println!("Line {}", idx+1);

        for (j, c) in line.chars().enumerate() {
            if c != '.' && c.is_digit(10) {
                temp = temp + c.to_string().as_str();
            }

            if temp.is_empty() && c != '.' && !c.is_alphanumeric() {
                starting_symbol = true;
            } else if starting_symbol && temp.is_empty() && c == '.' {
                starting_symbol = false;
            }

            if temp != "" && (c == '.' || j == line.len()-1) {
                if starting_symbol {
                    total += temp.parse::<i32>().unwrap();
                    println!("Adding: {}", temp);
                    temp = String::from("");
                    starting_symbol = false;
                    continue;
                }
                let l = std::cmp::max(0, j as isize - (temp.len()+1) as isize);

                if idx > 0 {
                    let above_line = lines.get(idx - 1).unwrap().get((l as usize)..j+1).unwrap();
                    let found = check_line(above_line);
                    if found {
                        println!("Adding: {}", temp);
                        total += temp.parse::<i32>().unwrap();
                        temp = String::from("");
                        continue;
                    }
                }

                if idx < lines.len()-1 {
                    let below_line = lines.get(idx + 1).unwrap().get((l as usize)..j+1).unwrap();
                    let found = check_line(below_line);
                    if found {
                        total += temp.parse::<i32>().unwrap();
                        println!("Adding: {}", temp);
                        temp = String::from("");
                        continue;
                    }
                }
                temp = String::from("");
            } else if temp != "" && !c.is_alphanumeric() {
                println!("Adding: {}", temp);
                total += temp.parse::<i32>().unwrap();
                temp = String::from("");
                starting_symbol = true;
            }
        }
    }

    println!("Total: {}", total);
}

fn part_2(filename: &str) {
    let mut total = 0;
    let input = std::fs::read_to_string(filename).unwrap();
    let lines: Vec<Vec<char>> = input.lines().map(|l| l.chars().collect()).collect();

    for (row, line) in lines.iter().enumerate() {
        for (col, c) in line.iter().enumerate() {
            if !is_symbol(*c) { continue; }

            let neighbors = [
                (row - 1, col - 1),
                (row - 1, col),
                (row - 1, col + 1),
                (row, col - 1),
                (row, col + 1),
                (row + 1, col - 1),
                (row + 1, col),
                (row + 1, col + 1),
            ];

            let found_gears: HashSet<usize> = neighbors
                .iter()
                .filter(|(row, col)| in_bounds(&lines, *row, *col))
                .flat_map(|(row, col)| adjacent_parts(&lines, *row, *col))
                .collect();

            if found_gears.len() == 2 {
                total += found_gears.iter().product::<usize>();
            }
        }
    }

    println!("Total: {}", total);
}

fn adjacent_parts(lines: &[Vec<char>], row: usize, col: usize) -> Option<usize> {
    if !lines[row][col].is_ascii_digit() { return None; }

    let mut start = col;
    while start > 0 && lines[row][start - 1].is_ascii_digit() {
        start -= 1;
    }


    let mut end = col;
    while end < lines[row].len()-1 && lines[row][end + 1].is_ascii_digit() {
        end += 1;
    }

    let num_string: String = lines[row][start..end+1].iter().collect();
    let num = num_string.parse().ok()?;

    return Some(num);
}

fn in_bounds(lines: &[Vec<char>], i: usize, j: usize) -> bool {
    lines.get(i).and_then(|l| l.get(j)).is_some()
}

fn is_symbol(c: char) -> bool {
    c == '*'
}

fn is_gear(c: char, adjacent: usize) -> bool {
    c == '*' && adjacent == 2
}

fn check_line(line: &str) -> bool {
    for c in line.chars() {
        if c != '.' && !c.is_alphanumeric() {
            return true;
        }
    }
    return false;
}

fn read_lines(filename: impl AsRef<Path>) -> Vec<String> {
    let file = File::open(filename).expect("no such file");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("Could not parse line"))
        .collect()
}

// Part 2 inspiration taken from: https://github.com/joao-conde/advents-of-code/blob/master/2023/src/bin/day03.rs
// Gave a nice new way to read in from files as well as a good idea of passing variables around in rust as well as its mapping functions.