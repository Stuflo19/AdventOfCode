use std::collections::HashMap;
use std::io::BufRead;

fn main() {
    part_1("input.txt");
    part_2("input.txt");
}

fn part_1(filename: &str) {
    let mut total = 0;
    let input = std::fs::read_to_string(filename).unwrap();
    let lines: Vec<&str> = input.lines().map(|l| l).collect();

    for (idx, line) in lines.iter().enumerate() {
        let mut num_matches = 0;
        let no_num: Vec<&str> = line.split(": ").collect();
        let cards = no_num.get(1).unwrap();
        let individual_cards: Vec<&str> = cards.split(" | ").collect();

        let winning_nums: Vec<&str> = individual_cards.get(0).unwrap().split(" ").filter(|c| *c != "").collect();
        let elfs_card: Vec<&str> = individual_cards.get(1).unwrap().split(" ").filter(|c| *c != "").collect();

        for num in elfs_card {
            if winning_nums.contains(&num) {
                num_matches += 1;
            }
        }

        let mut card_worth = 0;
        let total_matches = num_matches;
        while num_matches > 0 {
            num_matches -= 1;
            if total_matches - num_matches == 1 { card_worth += 1; }
            else { card_worth *= 2; }
        }

        total += card_worth;
    }

    println!("Part 1: {}", total);
}

fn part_2(filename: &str) {
    let mut collected_cards: HashMap<usize, i32> = HashMap::new();
    let input = std::fs::read_to_string(filename).unwrap();
    let lines: Vec<&str> = input.lines().map(|l| l).collect();

    for (idx, line) in lines.iter().enumerate() {
        let mut num_matches = 0;
        let no_num: Vec<&str> = line.split(": ").collect();
        let cards = no_num.get(1).unwrap();
        let individual_cards: Vec<&str> = cards.split(" | ").collect();

        let winning_nums: Vec<&str> = individual_cards.get(0).unwrap().split(" ").filter(|c| *c != "").collect();
        let elfs_card: Vec<&str> = individual_cards.get(1).unwrap().split(" ").filter(|c| *c != "").collect();

        for num in elfs_card {
            if winning_nums.contains(&num) {
                num_matches += 1;
            }
        }

        let mut num_current_card = 0;
        match collected_cards.get(&idx) {
            Some(val) => num_current_card = *val,
            None => {
                num_current_card = 1;
                collected_cards.insert(idx,  1);
            }
        }

        let mut cards_index = idx + num_matches;
        while cards_index > idx {
            let mut curr_val = 0;
            match collected_cards.get(&cards_index) {
                Some(val) => curr_val += (*val + num_current_card),
                None => curr_val = 1 + num_current_card
            }
            collected_cards.insert(cards_index, curr_val);
            cards_index -= 1;
        }
    }

    println!("Part 2: {}", collected_cards.values().fold(0, |sum, value| sum+value));
}