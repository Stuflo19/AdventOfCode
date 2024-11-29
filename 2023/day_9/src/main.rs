fn main() {
    let input = MirageMaintanence::from_file("input.txt");
    println!("Part 1: {:?}", input.part_1().unwrap());
    println!("Part 2: {:?}", input.part_2());
}

struct MirageMaintanence {
    input: Vec<Vec<isize>>,
}

impl MirageMaintanence {
    fn new(input: Vec<Vec<isize>>) -> Self {
        MirageMaintanence { input }
    }

    fn from_file(filename: &str) -> Self {
        let mut input: Vec<Vec<isize>> = Vec::new();
        let file = std::fs::read_to_string(filename).unwrap();
        for line in file.lines() {
            let mut line_vec: Vec<isize> = Vec::new();
            for num in line.split_whitespace() {
                line_vec.push(num.parse::<isize>().unwrap());
            }
            input.push(line_vec);
        }
        MirageMaintanence::new(input)
    }

    fn is_complete(&self, line: &[isize]) -> bool {
        line.iter().all(|&x| x == 0)
    }

    fn process_line(&self, line: &Vec<isize>, mut lines: Vec<Vec<isize>>) -> Vec<Vec<isize>> {
        let mut new_line = vec![];
        for (idx, num) in line.into_iter().enumerate() {
            if idx == 0 { continue; }
            new_line.push(num - line.get(idx - 1).unwrap());
        }

        if self.is_complete(&new_line) {
            return lines;
        }
        lines.push(new_line.clone());
        self.process_line(&new_line, lines)
    }

    fn calculate_sequence(&self, processed_lines: Vec<Vec<isize>>) -> Option<isize> {
        processed_lines.iter().rev().map(|line| line.last()).sum()
    }

    fn calculate_reverse_sequence(&self, processed_lines: Vec<Vec<isize>>) -> isize {
        let values = processed_lines.iter().rev().map(|line| line.first());
        values.fold(0, |acc, val| {
            val.unwrap_or(&0) - acc
        })
    }

    fn part_1(&self) -> Option<isize> {
        self.input.iter().map(|line| self.process_line(line, Vec::from([Vec::from(line.clone())]))).map(|processed_lines| self.calculate_sequence(processed_lines)).sum()
    }

    fn part_2(&self) -> isize {
        self.input.iter().map(|line| self.process_line(line, Vec::from([Vec::from(line.clone())]))).map(|processed_lines| self.calculate_reverse_sequence(processed_lines)).sum()
    }
}

#[cfg(test)]
mod tests {
    use crate::MirageMaintanence;

    #[test]
    fn part_1_test() {
        let input = MirageMaintanence::from_file("test.txt");
        let ans: Option<isize> = input.input.iter().map(|line| input.process_line(line, Vec::from([Vec::from(line.clone())]))).map(|processed_lines| input.calculate_sequence(processed_lines)).sum();
        assert_eq!(ans, Some(114));
    }

    #[test]
    fn from_file_test() {
        let input = MirageMaintanence::from_file("test.txt");
        assert_eq!(input.input, vec![vec![0, 3, 6, 9, 12, 15], vec![1, 3, 6, 10, 15, 21], vec![10, 13, 16, 21, 30, 45]]);
    }

    #[test]
    fn process_line_test() {
        let input = MirageMaintanence::from_file("test.txt");
        assert_eq!(input.process_line(&vec![0, 3, 6, 9, 12, 15], Vec::from([Vec::from([0, 3, 6, 9, 12, 15])])), vec![vec![0, 3, 6, 9, 12, 15], vec![3, 3, 3, 3, 3]]);
    }

    #[test]
    fn calculate_sequence_test() {
        let input = MirageMaintanence::from_file("test.txt");
        assert_eq!(input.calculate_sequence(vec![vec![0, 3, 6, 9, 12, 15], vec![3, 3, 3, 3, 3]]), Some(18));
    }

    #[test]
    fn reverse_sequence_test() {
        let input = MirageMaintanence::from_file("test.txt");
        assert_eq!(input.calculate_reverse_sequence(vec![vec![10, 13, 16, 21, 30, 45], vec![3, 3, 5, 9, 15], vec![0, 2, 4, 6], vec![2, 2, 2]]), 5);
    }
}
