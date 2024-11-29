use std::collections::HashMap;
use std::iter::Map;

fn main() {
    part_1();
    part_2();
}

fn part_1() {
    let paths = Paths::from_file("input.txt");
    let ans = paths.process_paths();
    println!("Part 1: {}", ans);
}

fn part_2() {
    let mut paths = Paths::from_file("input.txt");
    let p = paths.process_all_paths();
    println!("Part 2: {}", paths.lcm(p));
}

type Path = (String, String);

struct Paths {
    paths: HashMap<String, Path>,
    paths_with_a: Vec<String>,
    directions: String
}

impl Paths {
    fn new(paths: HashMap<String, Path>, directions: String, paths_with_a: Vec<String>) -> Self {
        Paths {
            paths,
            directions,
            paths_with_a
        }
    }

    fn from_file(filename: &str) -> Self {
        let input = std::fs::read_to_string(filename).unwrap();
        let mut lines: Vec<&str> = input.lines().map(|l| l).collect();
        let mut paths_with_a = Vec::new();

        let mut paths = HashMap::new();

        let directions = lines.remove(0).to_string();
        lines.remove(0);

        lines.into_iter().for_each(|l| {
            let split: Vec<&str> = l.split(" = ").collect();
            if split.get(0).unwrap().ends_with("A") { paths_with_a.push(split.get(0).unwrap().to_string()); }
            let path = split.get(1).unwrap().replace("(", "").replace(")", "").split(", ").map(|s| s.to_string()).collect::<Vec<String>>();
            paths.insert(split.get(0).unwrap().to_string(), (path.get(0).unwrap().to_string(), path.get(1).unwrap().to_string()));
        });

        Paths::new(paths, directions, paths_with_a)
    }

    fn add_path(&mut self, location: String, path: Path) {
        self.paths.insert(location, path);
    }

    fn get_location(&self, location: &str) -> Option<&Path> {
        self.paths.get(location)
    }

    fn take_path(&self, path: Path, direction: char) -> Path {
        match direction {
            'L' => self.get_location(&path.0),
            'R' => self.get_location(&path.1),
            _ => panic!("Invalid direction"),
        }.unwrap().clone()
    }

    fn get_path(&self, path: Path, direction: char) -> String {
        match direction {
            'R' => path.1,
            'L' => path.0,
            _ => panic!("Invalid direction"),
        }
    }

    fn process_paths_2(&mut self, mut path: Path) -> usize {
        let mut total = 1;
        let mut index = 0;

         loop {
             let direction = self.directions.chars().nth(index).unwrap();
             if self.get_path(path.clone(), direction).ends_with("Z") { break; }
             path = self.take_path(path, direction);
             total += 1;
             index += 1;
             if index == self.directions.len() { index = 0; }
        }

        total
    }

    fn process_all_paths(&mut self) -> Vec<usize> {
        Map::collect(self.paths_with_a.clone().iter().map(|e| self.process_paths_2(self.paths.get(e).unwrap().clone())))
    }

    fn lcm(&self, nums: Vec<usize>) -> usize {
        let mut result = 1;
        for i in nums {
            result = (i * result) / self.gcd(i, result);
        }
        result
    }

    fn gcd(&self, a: usize, b: usize) -> usize {
        if b == 0 { return a; }
        self.gcd(b, a % b)
    }

    fn process_paths(&self) -> usize {
        let mut path = self.paths.get("AAA").unwrap().clone();
        let mut total = 1;
        let mut index = 0;

        loop {
            let direction = self.directions.chars().nth(index).unwrap();
            if direction == 'L' && path.0 == "ZZZ" { break; }
            if direction == 'R' && path.1 == "ZZZ" { break; }
            path = self.take_path(path, direction);
            total += 1;
            index += 1;
            if index == self.directions.len() { index = 0; }
        }

        total
    }
}

#[cfg(test)]
mod tests {
    use std::collections::HashMap;
    use crate::{Path, Paths};

    #[test]
    fn part_1_test() {
        let mut paths = Paths::new(HashMap::new(), String::from("LLR"), Vec::new());
        paths.add_path("AAA".to_string(), ("BBB".to_string(), "BBB".to_string()));
        paths.add_path("BBB".to_string(), ("AAA".to_string(), "ZZZ".to_string()));
        paths.add_path("ZZZ".to_string(), ("ZZZ".to_string(), "ZZZ".to_string()));

        let ans = paths.process_paths();
        assert_eq!(ans, 6);
    }

    #[test]
    fn part_2_test() {
        let mut paths = Paths::from_file("test3.txt");
        let p = paths.process_all_paths();

        assert_eq!(paths.lcm(p), 6);
    }

    #[test]
    fn initial_test() {
        let mut paths = Paths::new(HashMap::new(), String::from("RL"), Vec::new());
        paths.add_path("AAA".to_string(), ("BBB".to_string(), "CCC".to_string()));
        paths.add_path("BBB".to_string(), ("DDD".to_string(), "EEE".to_string()));
        paths.add_path("CCC".to_string(), ("ZZZ".to_string(), "GGG".to_string()));
        paths.add_path("DDD".to_string(), ("DDD".to_string(), "DDD".to_string()));
        paths.add_path("EEE".to_string(), ("EEE".to_string(), "EEE".to_string()));
        paths.add_path("GGG".to_string(), ("GGG".to_string(), "GGG".to_string()));
        paths.add_path("ZZZ".to_string(), ("ZZZ".to_string(), "ZZZ".to_string()));

        let ans = paths.process_paths();
        assert_eq!(ans, 2);
    }

    #[test]
    fn test_take_path() {
        let path: Path = ("AAA".to_string(), "BBB".to_string());
        let mut paths = Paths::new(HashMap::new(), String::from("RL"), Vec::new());
        paths.add_path("AAA".to_string(), path.clone());

        assert_eq!(paths.take_path(path.clone(), 'L'), path.clone());
    }

    #[test]
    fn from_file_test() {
        let paths = Paths::from_file("test.txt");
        let test1 = paths.process_paths();
        assert_eq!(test1, 2);

        let paths2 = Paths::from_file("test2.txt");
        let test2 = paths2.process_paths();
        assert_eq!(test2, 6);
    }
}