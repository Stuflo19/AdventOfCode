mod part_1;
mod part_2;

use std::collections::HashMap;
use itertools::sorted;
use crate::part_1::part_1;
use crate::part_2::part_2;

fn main() {
    part_1();
    part_2();
}

const CARD_SCORES: [char; 13] = ['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'];
const CARD_SCORES_JOKER: [char; 13] = ['A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'];

fn get_card_score(card: char, apply_jokers: bool) -> Option<usize> {
    match apply_jokers {
        true => CARD_SCORES_JOKER.iter().position(|&e| e == card),
        false => CARD_SCORES.iter().position(|&e| e == card)
    }
}

type Hand = (String, usize);
type HandWorth = (usize, usize);
#[derive(Debug)]
struct CamelCards {}

struct CardRankings {
    game_scores: Vec<Hand>,
    camel_cards: CamelCards
}

impl CardRankings {
    fn new() -> Self {
        let game_scores: Vec<Hand> = Vec::new();
        let camel_cards = CamelCards::new();
        Self { game_scores, camel_cards }
    }

    fn add_hand(&mut self, hand: Hand, apply_jokers: bool) {
        let len = self.game_scores.len();
        if len == 0 {
            self.game_scores.push(hand);
            return;
        }

        for (idx, h) in self.game_scores.iter().enumerate() {
            if !self.camel_cards.compare_hands(&h.0, &hand.0, apply_jokers) {
                self.game_scores.insert(idx, hand);
                return;
            }
        }

        if len == self.game_scores.len() {
            self.game_scores.push(hand);
        }
    }

    fn total_score(&self) -> usize {
        self.game_scores.iter().enumerate().map(|(idx, e)| e.1 * (idx+1)).sum()
    }
}

impl CamelCards {
    fn new() -> Self {
        Self{ }
    }

    fn compare_equal_hands(&self, hand_1: &str, hand_2: &str, apply_jokers: bool) -> bool {
        let hand_1_chars: Vec<char> = hand_1.chars().collect();
        let hand_2_chars: Vec<char> = hand_2.chars().collect();

        for i in 0..hand_1_chars.len() {
            let hand_1_score = get_card_score(hand_1_chars[i], apply_jokers).unwrap();
            let hand_2_score = get_card_score(hand_2_chars[i], apply_jokers).unwrap();

            if hand_1_score == hand_2_score { continue; }
            else if hand_1_score < hand_2_score { return true; }
            else { return false; }
        }
        return false;
    }

    fn compare_hands(&self, hand_1: &str, hand_2: &str, apply_jokers: bool) -> bool {
        let hand_1_worth = self.check_hand(hand_1, apply_jokers);
        let hand_2_worth = self.check_hand(hand_2, apply_jokers);

        return
        if hand_1_worth == hand_2_worth { !self.compare_equal_hands(hand_1, hand_2, apply_jokers) }
        else if hand_1_worth < hand_2_worth { true }
        else { false }
    }

    // Part 2 check hand:
    // Sorting for counting jacks inspiration from: https://advent-of-code.xavd.id/writeups/2023/day/7/
    fn check_hand(&self, hand: &str, apply_jokers: bool) -> usize {
        const FACTOR: usize = usize::pow(13, 5);

        let chars: Vec<char> = hand.chars().collect();
        let mut char_map: HashMap<char, usize> = HashMap::new();

        chars.into_iter().for_each(|e| {
            char_map.insert(e, match char_map.get(&e) {
                Some(value) => value + 1,
                None => 1
            });
        });

        if apply_jokers {
            let num_jokers = *char_map.get(&'J').unwrap_or(&0);
            if num_jokers < 5 {
                char_map.remove(&'J');
                let max = char_map.iter().max_by(|a, b| a.1.cmp(b.1)).unwrap();
                char_map.insert(*max.0, max.1 + num_jokers);
            }
        }

        return match sorted(char_map.values()).as_slice() {
            [5] => 6,
            [1, 4] => 5,
            [2, 3] => 4,
            [1, 1, 3] => 3,
            [1, 2, 2] => 2,
            [1, 1, 1, 2] => 1,
            [1, 1, 1, 1, 1] => 0,
            _ => { println!("Error: {:?}", char_map); 0 }
        }

        // let values: Vec<&usize> = char_map.iter().filter(|(&key, &val)| key != 'J').map(|(key, val)| val).collect();
        // let num_jokers = *char_map.get(&'J').unwrap_or(&0);
        // let mut highest = 0;
        // let mut second_highest= 0;
        //
        // // println!("{:?}", values);
        // for count in values.iter() {
        //     if **count > highest {
        //         second_highest = highest;
        //         highest = **count;
        //     } else if **count > second_highest {
        //         second_highest = **count;
        //     }
        // }
        //
        // let highest_jokers = highest + num_jokers;
        //
        // let strength =
        //     if highest_jokers == 5 { 6 }
        //     else if highest_jokers == 4 { 5 }
        //     else if highest_jokers == 3 && second_highest == 2 { 4 }
        //     else if highest_jokers == 3 && second_highest == 1 { 3 }
        //     else if highest_jokers == 2 && second_highest == 2 { 2 }
        //     else if highest_jokers == 2 && second_highest == 1 { 1 }
        //     else { 0 };
        //
        // return strength;
    }
}

