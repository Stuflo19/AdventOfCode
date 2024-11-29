use crate::{CardRankings};

pub fn part_1() {
    let input = std::fs::read_to_string("input.txt").unwrap();
    let lines: Vec<&str> = input.lines().map(|l| l).collect();

    let mut ranking = CardRankings::new();

    lines.into_iter().for_each(|l| {
        let mut split = l.split_whitespace();
        let hand = split.next().unwrap();
        let score = split.next().unwrap().parse::<usize>().unwrap();

        ranking.add_hand((String::from(hand), score), false);
    });

    println!("Part 1: {}", ranking.total_score());
}

#[cfg(test)]
mod tests {
    use crate::{CardRankings};

    #[test]
    fn part_1_test() {
        let mut ranking = CardRankings::new();

        ranking.add_hand((String::from("32T3K"), 765), false);
        ranking.add_hand((String::from("T55J5"), 684), false);
        ranking.add_hand((String::from("KK677"), 28), false);
        ranking.add_hand((String::from("KTJJT"), 220), false);
        ranking.add_hand((String::from("QQQJA"), 483), false);

        // println!("{:?}", ranking.game_scores);

        assert_eq!(ranking.total_score(), 6440)
    }

    // #[test]
    // fn five_of_kind_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("AAAAA"), 6)
    // }
    //
    // #[test]
    // fn four_of_kind_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("AA8AA"), 5)
    // }
    //
    // #[test]
    // fn full_house_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("23332"), 4)
    // }
    //
    // #[test]
    // fn three_of_kind_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("TTT98"), 3)
    // }
    //
    // #[test]
    // fn two_pair_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("23432"), 2)
    // }
    //
    // #[test]
    // fn one_pair_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("A23A4"), 1)
    // }
    //
    // #[test]
    // fn high_card_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("23456"), 0)
    // }

    // #[test]
    // fn compare_hands_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.compare_hands("KKKKK", "AAAAA"), false)
    // }
    //
    // #[test]
    // fn compare_hands_test2() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.compare_hands("KK677", "KTJJT"), true)
    // }
}