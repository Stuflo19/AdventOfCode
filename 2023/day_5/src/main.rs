use std::fmt::Error;
use std::ops::Range;

fn main() {
    let input = std::fs::read_to_string("test.txt").unwrap();
    // let input = std::fs::read_to_string("input.txt").unwrap();
    let lines: Vec<&str> = input.lines().map(|l| l).collect();

    let almanac: Almanac = Almanac::parse(lines).unwrap();

    let p1 = almanac.p1();
    println!("Part 1: {}", p1);

    let p2 = almanac.p2();
    println!("Part 2: {}", p2);
}

#[derive(Debug)]
struct Mapping {
    range: Range<isize>,
    delta: isize,
    source: isize,
    destination: isize,
    length: isize
}

impl Mapping {
    fn new(source: isize, destination: isize, length: isize) -> Self {
        Self {
            range: source..source + length,
            delta: destination - source,
            source,
            destination,
            length
        }
    }

    fn contains(&self, value: isize) -> bool {
        self.range.contains(&value)
    }

    fn apply(&self, value: isize) -> isize {
        value + self.delta
    }
}

#[derive(Debug)]
struct Almanac {
    seeds: Vec<isize>,
    mappings: Vec<Vec<Mapping>>,
}

impl Almanac {
    fn new(seeds: Vec<isize>, mappings: Vec<Vec<Mapping>>) -> Self {
        Self { seeds, mappings }
    }

    fn apply_seed(&self, seed: isize) -> isize {
        let mut value = seed;
        for mapping in &self.mappings {
            'range: for m in mapping {
                if m.contains(value) {
                    value = m.apply(value);
                    break 'range;
                }
            }
        }
        value
    }

    // Part 1 inspiration from: https://gist.github.com/icub3d/1d21197f4b6eaabbdf0a43fd6a25ba1a
    fn p1(&self) -> isize {
        self.seeds
            .iter()
            .map(|s| self.apply_seed(*s))
            .min()
            .unwrap()
    }

    // Part 2 from: https://gist.github.com/HakanVardarr/5dc98fd1fc0e7897136280467dad6991
    fn p2(&self) -> isize {
        let mut reverse = 1;
        loop {
            let mut curent_reverse = reverse;
            'next_map: for i in (0..self.mappings.len()).rev() {
                let mut changed = false;
                for s in &self.mappings[i] {
                    if changed {
                        continue 'next_map;
                    }
                    let i = s.destination;
                    let k = i + s.length - 1;

                    if curent_reverse >= i && curent_reverse <= k {
                        curent_reverse = curent_reverse + s.source - s.destination;
                        changed = true
                    }
                }
            }
            for seeds in self.seeds.chunks(2) {
                if curent_reverse >= seeds[0] && curent_reverse < seeds[1] + seeds[0] {
                    return reverse
                }
            }
            reverse += 1;
        }
    }

    fn parse(input: Vec<&str>) -> Result<Almanac, Error> {
        let mut seeds: Vec<isize> = vec![];
        let mut mappings: Vec<Vec<Mapping>> = vec![];
        let mut mapping: Vec<Mapping> = vec![];
        for (idx, line) in input.iter().enumerate() {
            if idx == 0 {
                let split: Vec<&str> = line.split(": ").collect();
                seeds = split.get(1).unwrap().split_whitespace().map(|e| e.parse::<isize>().unwrap()).collect();
            }

            if line == &"" { continue; }
            else if !line.chars().next().unwrap().is_digit(10) {
                if mapping.is_empty() { continue; }
                mappings.push(mapping);
                mapping = vec![];
            }
            else {
                let values: Vec<isize> = line.split_whitespace().map(|e| e.parse::<isize>().unwrap()).collect();
                mapping.push(Mapping::new(values[1], values[0], values[2]))
            }
        }
        if !mapping.is_empty() { mappings.push(mapping) }

        Ok(Self::new(seeds, mappings))
    }
}

// Too low: 303894139