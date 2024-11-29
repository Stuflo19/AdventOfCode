use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    // part_1("input.txt");
    part_2("input.txt");
}

fn part_1(filename: &str) {
    let mut total = 0;
    let mut game_num = 0;
    let mut raw_total = 0;

    if let Ok(lines) = read_lines(filename) {
        let mut max_colours: HashMap<&str, i32> = HashMap::new();

        for line in lines {
            if let Ok(input) = line {
                game_num += 1;
                raw_total += game_num;

                let split: Vec<&str> = input.split(": ").collect();
                let curr_game: Vec<&str> = split[1].split(";").collect();
                for hands in curr_game {
                    let mut impossible = false;
                    let hand: Vec<&str> = hands.split(",").collect();
                    for colours in hand {
                        let colour: Vec<&str> = colours.split_whitespace().collect();
                        if colour[0].parse::<i32>().unwrap() > *max_colours.get(colour[1]).unwrap() {
                            total += game_num;
                            impossible = true;
                            break;
                        }
                    }

                    if impossible { break; }
                }
            }
        }

        println!("Total: {}", raw_total - total);
    }
}

fn part_2(filename: &str) {
    let mut total = 0;
    let mut game_num = 0;

    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(input) = line {
                game_num += 1;

                let split: Vec<&str> = input.split(": ").collect();
                let curr_game: Vec<&str> = split[1].split(";").collect();

                let mut max_red = 0;
                let mut max_blue = 0;
                let mut max_green = 0;

                for hands in curr_game {
                    let hand: Vec<&str> = hands.split(",").collect();

                    for colours in hand {
                        let colour: Vec<&str> = colours.split_whitespace().collect();

                        match colour[1] {
                            "blue"=> {
                                let value = colour[0].parse::<i32>().unwrap();
                                if value > max_blue {
                                    max_blue = value
                                }
                            }
                            "red"=> {
                                let value = colour[0].parse::<i32>().unwrap();
                                if value > max_red {
                                    max_red = value
                                }
                            }
                            "green"=> {
                                let value = colour[0].parse::<i32>().unwrap();
                                if value > max_green {
                                    max_green = value
                                }
                            }
                            _ => {}
                        }
                    }
                }
                println!("Game: {}, Total: {}", game_num, (max_green * max_blue * max_red));
                total += (max_green * max_blue * max_red)
            }
        }

        println!("Total: {}", total);
    }
}
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
