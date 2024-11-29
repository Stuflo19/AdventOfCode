fn main() {
    println!("Part 1: {}", part_1());
    println!("Part 2: {}", part_2());
}

fn part_1() -> usize {
    let race_1 = BoatRace::new(643, 61);
    let race_2 = BoatRace::new(1184, 70);
    let race_3 = BoatRace::new(1362, 90);
    let race_4 = BoatRace::new(1041, 66);

    race_1.get_possible_wins() * race_2.get_possible_wins() * race_3.get_possible_wins() * race_4.get_possible_wins()
}


fn part_2() -> usize {
    let race = BoatRace::new(643118413621041, 61709066);

    race.get_possible_wins()
}

#[derive(Debug)]
struct BoatRace {
    distance: usize,
    duration: usize
}

impl BoatRace {
    fn new(distance: usize, duration: usize) -> Self {
        Self { distance, duration }
    }

    fn get_possible_wins(&self) -> usize {
        (0..self.duration).filter(|e| (e * (self.duration - e)) > self.distance).count()
    }
}

#[cfg(test)]
mod tests {
    use crate::BoatRace;

    #[test]
    fn boat_timing_calculation() {
        let race = BoatRace::new(9, 7);

        assert_eq!(race.get_possible_wins(), 4)
    }

    #[test]
    fn part_1_test() {
        let race_1 = BoatRace::new(9, 7);
        let race_2 = BoatRace::new(40, 15);
        let race_3 = BoatRace::new(200, 30);

        let result = race_1.get_possible_wins() * race_2.get_possible_wins() * race_3.get_possible_wins();

        assert_eq!(result, 288)
    }

    #[test]
    fn part_2_test() {
        let race = BoatRace::new(940200, 71530);

        let result = race.get_possible_wins();

        assert_eq!(result, 71503)
    }
}
