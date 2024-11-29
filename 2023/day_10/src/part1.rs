const INPUT: &'static str = include_str!("../input.txt");
const TEST_INPUT: &'static str = include_str!("../test.txt");
const TEST_INPUT2: &'static str = include_str!("../test2.txt");
const TEST2_INPUT: &'static str = include_str!("../test_part2.txt");
const TEST2_INPUT2: &'static str = include_str!("../test2_part2.txt");
const TEST2_INPUT3: &'static str = include_str!("../test3_part2.txt");

#[derive(Clone, Debug, PartialEq)]
enum Direction {
    North,
    South,
    East,
    West,
    NS,
    EW,
    NE,
    NW,
    SW,
    SE,
    Ground,
    Start
}

#[derive(Clone, Debug, PartialEq, Copy)]
struct Coordinate {
    x: usize,
    y: usize
}

#[derive(Clone, Debug)]
struct Node {
    neighbours: Vec<Coordinate>,
    direction: Direction,
}

struct Graph {
    start_node: Coordinate,
    vertices: Vec<Vec<Node>>,
    location: Coordinate,
}

impl Graph {
    fn from_file(contents: &'static str) -> Self {
        let mut graph: Vec<Vec<Node>> = Vec::new();
        let mut start_node: Coordinate = Coordinate { x: 0, y: 0 };
        let height = contents.lines().count();
        let width = contents.lines().nth(0).unwrap().len();

        for (y, line) in contents.lines().enumerate() {
            graph.push(Vec::new());
            for (x, c) in line.chars().enumerate() {
                let mut neighbours = vec![];
                match c {
                    '|' => {
                        if y + 1 < height { neighbours.push(Coordinate { x, y: y+1 }) }
                        if y > 0 { neighbours.push(Coordinate { x, y: y-1 }) }
                        graph[y].push(Node{ neighbours, direction: Direction::NS });
                    }
                    '-' => {
                        if x + 1 < width { neighbours.push(Coordinate { x: x+1, y }) }
                        if x > 0 { neighbours.push(Coordinate { x: x-1, y }) }
                        graph[y].push(Node{ neighbours, direction: Direction::EW });
                    }
                    'L' => {
                        if x < width { neighbours.push(Coordinate { x: x+1, y }) }
                        if y > 0 { neighbours.push(Coordinate { x, y: y-1 }) }
                        graph[y].push(Node{ neighbours, direction: Direction::NE });
                    }
                    'J' => {
                        if x > 0 { neighbours.push(Coordinate { x: x-1, y }) }
                        if y > 0 { neighbours.push(Coordinate { x, y: y-1 }) }
                        graph[y].push(Node{ neighbours, direction: Direction::NW });
                    }
                    '7' => {
                        if x > 0 { neighbours.push(Coordinate { x: x-1, y }) }
                        if y + 1 < height { neighbours.push(Coordinate { x, y: y+1 }) }
                        graph[y].push(Node{ neighbours, direction: Direction::SW });
                    }
                    'F' => {
                        if x < width { neighbours.push(Coordinate { x: x+1, y }) }
                        if y < height { neighbours.push(Coordinate { x, y: y+1 }) }
                        graph[y].push(Node{ neighbours, direction: Direction::SE });
                    }
                    '.' => {
                        graph[y].push(Node{ neighbours: Vec::new(), direction: Direction::Ground });
                    }
                    'S' => {
                        if x < width { neighbours.push(Coordinate { x: x+1, y }) }
                        if y < height { neighbours.push(Coordinate { x, y: y+1 }) }
                        if x > 0 { neighbours.push(Coordinate { x: x-1, y }) }
                        if y > 0 { neighbours.push(Coordinate { x, y: y-1 }) }
                        start_node = Coordinate { x, y };
                        graph[y].push(Node{ neighbours, direction: Direction::Start });
                    }
                    _ => { println!("Unimplemented Character: {}", c) }
                }
            }
        }

        Self {
            start_node,
            vertices: graph,
            location: Coordinate { x: 0, y: 0 }
        }
    }

    fn direction_moved (&self, from: &Coordinate, to: &Coordinate) -> Direction {
        if from.x == to.x {
            if from.y > to.y {
                Direction::North
            } else {
                Direction::South
            }
        } else {
            if from.x > to.x {
                Direction::West
            } else {
                Direction::East
            }
        }
    }

    fn walk(&self) -> Vec<Coordinate> {
        let mut pipe_loop: Vec<Coordinate> = vec![self.start_node];
        let start= self.vertices.get(self.start_node.y).unwrap().get(self.start_node.x).unwrap().neighbours.iter().filter(|n| {
            let node = self.vertices.get(n.y).unwrap().get(n.x).unwrap();
            node.neighbours.contains(&self.start_node)
        }).collect::<Vec<&Coordinate>>().get(0).unwrap().clone();

        let mut from = self.direction_moved(&self.start_node, &start);
        let mut current = start.clone();
        while current != self.start_node {
            pipe_loop.push(current.clone());
            let node = self.vertices.get(current.y).unwrap().get(current.x).unwrap();
            let next = node.neighbours.iter().filter(|n| {
                let nod = self.vertices.get(n.y).unwrap().get(n.x).unwrap();
                nod.neighbours.contains(&current) && self.direction_moved(n, &current) != from
            }).collect::<Vec<&Coordinate>>();
            let n = next.get(0).unwrap().clone().clone();
            from = self.direction_moved(&current, &n);
            current = n;
        }
        pipe_loop.push(current);

        // println!("Pipe Loop: {:?}", pipe_loop);
        pipe_loop
    }

    fn shoelace(&self, pipe_loop: Vec<Coordinate>) -> isize {
        let a = pipe_loop.windows(2).map(|w| {
            (w[0].x * w[1].y) as isize - (w[0].y * w[1].x) as isize
        }).sum::<isize>();
        (a.abs() / 2) - (pipe_loop.len() as isize / 2) + 1
    }
}

pub fn part_1() {
    let graph = Graph::from_file(INPUT);
    let result = graph.walk().len() / 2;
    println!("Part 1: {}", result);
}

pub fn part_2() {
    let graph = Graph::from_file(INPUT);
    println!("Part 2: {}", graph.shoelace(graph.walk()));
}

#[cfg(test)]
mod tests {
    use crate::part1::{Graph, TEST2_INPUT, TEST2_INPUT2, TEST2_INPUT3, TEST_INPUT2};
    use crate::part1::TEST_INPUT;

    #[test]
    fn part_1_test() {
        let graph = Graph::from_file(TEST_INPUT);
        let res = graph.walk().len();

        assert_eq!(res, 8);
    }

    #[test]
    fn part_1_test2() {
        let graph = Graph::from_file(TEST_INPUT2);
        let res = graph.walk().len();

        assert_eq!(res, 16);
    }

    #[test]
    fn part_2_test() {
        let graph = Graph::from_file(TEST2_INPUT);
        let res = graph.shoelace(graph.walk());

        assert_eq!(res, 4);
    }
    #[test]
    fn part_2_test2() {
        let graph = Graph::from_file(TEST2_INPUT2);
        let res = graph.shoelace(graph.walk());

        assert_eq!(res, 8);
    }
    #[test]
    fn part_2_test3() {
        let graph = Graph::from_file(TEST2_INPUT3);
        let res = graph.shoelace(graph.walk());

        assert_eq!(res, 10);
    }
}