const INPUT: &'static str = include_str!("../input.txt");
const TEST: &'static str = include_str!("../test.txt");

fn main() {
    // let galaxy = Galaxy::from_file(INPUT);
    // let result = galaxy.part_1();
    let result = Galaxy::part_1();
    println!("Part 1: {}", result);

    let result = Galaxy::part_2();
    println!("Part 2: {}", result);
}

#[derive(Debug, PartialEq)]
struct Coordinate {
    x: i64,
    y: i64,
}

impl Coordinate {
    fn new(x: i64, y: i64) -> Coordinate {
        Coordinate { x, y }
    }

    fn distance_between(&self, other: &Coordinate) -> i64 {
        (self.x - other.x).abs() + (self.y - other.y).abs()
    }
}

#[derive(Debug)]
struct Galaxy {
    stars: Vec<Coordinate>,
}

impl Galaxy {
    fn new(stars: Vec<Coordinate>) -> Self {
        Self {
            stars,
        }
    }

    fn from_file(contents: &'static str) -> Self {
        let mut stars = Vec::new();
        let galaxy_vec = expand_input(contents);

        for (row, row_vec) in galaxy_vec.iter().enumerate() {
            for (col, char) in row_vec.iter().enumerate() {
                if *char == '#' { stars.push(Coordinate::new(col as i64, row as i64)); }
            }
        }

        Galaxy::new(stars)
    }

    fn from_file2(contents: &'static str, expansion_factor: usize) -> Self {
        let mut stars = Vec::new();
        let (new_rows, new_columns) = expand_input2(contents);

        for (row, row_vec) in contents.lines().enumerate() {
            for (col, char) in row_vec.chars().enumerate() {
                let col_expanded = col + (new_columns.iter().filter(|e| *e < &col).count() * expansion_factor);
                let row_expanded = row + (new_rows.iter().filter(|e| *e < &row).count() * expansion_factor);
                if char == '#' {
                    stars.push(Coordinate::new(col_expanded as i64, row_expanded as i64));
                }
            }
        }

        Galaxy::new(stars)
    }

    fn part_1() -> i64 {
        let galaxy = Galaxy::from_file(INPUT);
        galaxy.distance_between_pairs().iter().sum::<i64>()
    }

    fn part_2() -> i64 {
        let galaxy = Galaxy::from_file2(INPUT, 999999);
        galaxy.distance_between_pairs().iter().sum::<i64>()
    }

    fn distance_between_pairs(&self) -> Vec<i64> {
        let mut distances = Vec::new();
        for (idx, star) in self.stars.iter().enumerate() {
            for other_star in &self.stars[idx + 1..] {
                distances.push(star.distance_between(other_star));
            }
        }
        distances
    }
}

type Expansions = (Vec<usize>, Vec<usize>);

fn expand_input(contents: &'static str) -> Vec<Vec<char>> {
    let mut galaxy_vec: Vec<Vec<char>> = Vec::new();
    let mut new_rows: Vec<usize> = Vec::new();
    let mut new_columns: Vec<usize> = Vec::new();
    for line in contents.lines() {
        galaxy_vec.push(line.chars().collect());
    }

    for (col, row) in galaxy_vec.iter().enumerate() {
        if requires_expansion(row) { new_rows.push(col); }
        if requires_expansion(&galaxy_vec.iter().map(|e| e[col]).collect::<Vec<char>>()) {
            new_columns.push(col);
        }
    }

    for (idx, row) in new_rows.iter().enumerate() {
        galaxy_vec.insert(*row + idx, vec!['.'; galaxy_vec[0].len()]);
    }

    for (idx, col) in new_columns.iter().enumerate() {
        galaxy_vec.iter_mut().for_each(|e| e.insert(*col + idx, '.'));
    }
    galaxy_vec
}

fn expand_input2(contents: &'static str) -> Expansions {
    let mut new_rows: Vec<usize> = Vec::new();
    let mut new_columns: Vec<usize> = Vec::new();
    let galaxy_vec: Vec<Vec<char>> = contents.lines().map(|e| e.chars().collect::<Vec<char>>()).collect();

    for (col, row) in galaxy_vec.iter().enumerate() {
        if requires_expansion(row) { new_rows.push(col); }
        if requires_expansion(&galaxy_vec.iter().map(|e| e[col]).collect::<Vec<char>>()) {
            new_columns.push(col);
        }
    }

    (new_rows, new_columns)
}

fn requires_expansion(test: &[char]) -> bool {
    test.iter().all(|e| *e == '.')
}

#[cfg(test)]
mod tests {
    use crate::{Coordinate, expand_input, Galaxy, TEST};

    #[test]
    fn test_distance_between() {
        let galaxy = Galaxy::from_file2(TEST, 1);
        assert_eq!(galaxy.stars.get(4).unwrap().distance_between(galaxy.stars.get(8).unwrap()), 9);
        assert_eq!(galaxy.stars.get(0).unwrap().distance_between(galaxy.stars.get(6).unwrap()), 15);
        assert_eq!(galaxy.stars.get(2).unwrap().distance_between(galaxy.stars.get(5).unwrap()), 17);
        assert_eq!(galaxy.stars.get(7).unwrap().distance_between(galaxy.stars.get(8).unwrap()), 5);
    }

    #[test]
    fn test_expand_input() {
        let expanded_input =
"....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......";
        let expanded_vec = expanded_input.lines().map(|e| e.chars().collect::<Vec<char>>()).collect::<Vec<Vec<char>>>();
        let galaxy_vec = expand_input(TEST);

        assert_eq!(galaxy_vec, expanded_vec);
    }

    #[test]
    fn test_stars_from_file() {
        let expanded_input =
"....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......";

        let mut test_stars = Vec::new();
        let galaxy = Galaxy::from_file(TEST);

        for (row, line) in expanded_input.lines().enumerate() {
            for (col, char) in line.chars().enumerate() {
                if char == '#' { test_stars.push(Coordinate::new(col as i64, row as i64)) }
            }
        };

        assert_eq!(galaxy.stars, test_stars);
    }

    #[test]
    fn part_1_test() {
        let galaxy = Galaxy::from_file2(TEST, 1);
        let result: i64 = galaxy.distance_between_pairs().iter().sum();

        assert_eq!(result, 374);
    }

    #[test]
    fn part_2_test() {
        let galaxy = Galaxy::from_file2(TEST, 9);
        let result: i64 = galaxy.distance_between_pairs().iter().sum();

        assert_eq!(result, 1030);
    }
    #[test]
    fn part_2_test2() {
        let galaxy = Galaxy::from_file2(TEST, 99);
        let result: i64 = galaxy.distance_between_pairs().iter().sum();

        assert_eq!(result, 8410);
    }
}
