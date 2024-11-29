use crate::{CardRankings};

pub fn part_2() {
    // 253499763
    let input = std::fs::read_to_string("input.txt").unwrap();
    let lines: Vec<&str> = input.lines().map(|l| l).collect();

    let mut ranking = CardRankings::new();

    lines.into_iter().for_each(|l| {
        let mut split = l.split_whitespace();
        let hand = split.next().unwrap();
        let score = split.next().unwrap().parse::<usize>().unwrap();

        ranking.add_hand((String::from(hand), score), true);
    });

    // println!("{:?}", ranking.game_scores);

    println!("Part 2: {}", ranking.total_score());
}

#[cfg(test)]
mod tests {
    use crate::{CardRankings};

    #[test]
    fn part_2_test() {
        let mut ranking = CardRankings::new();

        ranking.add_hand((String::from("32T3K"), 765), true);
        ranking.add_hand((String::from("T55J5"), 684), true);
        ranking.add_hand((String::from("KK677"), 28), true);
        ranking.add_hand((String::from("KTJJT"), 220), true);
        ranking.add_hand((String::from("QQQJA"), 483), true);


        println!("{:?}", ranking.game_scores);

        assert_eq!(ranking.total_score(), 5905)
    }

    // #[test]
    // fn five_of_kind_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("8JJJJ"), 6);
    // }
    //
    // #[test]
    // fn four_of_kind_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("QQQJA"), 5);
    //     assert_eq!(cc.check_hand("T55J5"), 5);
    //     assert_eq!(cc.check_hand("KTJJT"), 5);
    // }
    //
    // #[test]
    // fn full_house_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("233J2"), 4);
    //     assert_eq!(cc.check_hand("KKJ77"), 4);
    // }
    //
    // #[test]
    // fn three_pair_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("6JAJ7"), 3);
    // }
    //
    // #[test]
    // fn two_pair_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("22334"), 2);
    // }
    //
    // #[test]
    // fn one_pair_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("22345"), 1);
    //     assert_eq!(cc.check_hand("J2345"), 1);
    // }
    //
    // #[test]
    // fn high_card_test() {
    //     let cc = CamelCards::new();
    //
    //     assert_eq!(cc.check_hand("23456"), 0);
    // }
}