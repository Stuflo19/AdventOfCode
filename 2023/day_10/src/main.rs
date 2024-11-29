mod part1;
use part1::part_1;

use std::collections::{VecDeque, HashMap, HashSet};
use crate::part1::part_2;

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
struct Coordinate {
    x: usize,
    y: usize,
}

struct Graph {
    start: Coordinate,
    vertices: Vec<Coordinate>,
    adjacency_list: Vec<Vec<usize>>,
}

impl Graph {
    fn from_file(filename: &str) -> Self {
        let contents = std::fs::read_to_string(filename).unwrap();
        let mut coordinates = Vec::new();

        for (y, line) in contents.lines().enumerate() {
            for (x, _c) in line.chars().enumerate() {
                coordinates.push(Coordinate { x, y });
            }
        }

        let mut graph = Graph::new(coordinates.clone(), Coordinate { x: 0, y: 0 });

        for (y, line) in contents.lines().enumerate() {
            for (x, c) in line.chars().enumerate() {
                graph.add_edges(x, y, c, line.len(), contents.lines().count());
            }
        }

        graph
    }
    fn add_edges(&mut self, x: usize, y: usize, char: char, height: usize, width: usize)  {
        match char {
            // North and South
            '|' => {
                if y + 1 < height { self.add_edge(Coordinate { x, y }, Coordinate { x, y: y+1 }); }
                if y >= 1 { self.add_edge(Coordinate { x, y }, Coordinate { x, y: y-1 }); }
            }
            // West and East
            '-' => {
                if x + 1 < width { self.add_edge(Coordinate { x, y }, Coordinate { x: x+1, y }); }
                if x >= 1 { self.add_edge(Coordinate { x, y }, Coordinate { x: x-1, y }); }
            }
            // North and east
            'L' => {
                if x + 1 < width { self.add_edge(Coordinate { x, y }, Coordinate { x: x+1, y }); }
                if y >= 1 { self.add_edge(Coordinate { x, y }, Coordinate { x, y: y-1 }); }
            }
            // North and west
            'J' => {
                if x >= 1 { self.add_edge(Coordinate { x, y }, Coordinate { x: x-1, y }); }
                if y >= 1 { self.add_edge(Coordinate { x, y }, Coordinate { x, y: y-1 }); }
            }
            // South and west
            '7' => {
                if x >= 1 { self.add_edge(Coordinate { x, y }, Coordinate { x: x-1, y }); }
                if y + 1 < height { self.add_edge(Coordinate { x, y }, Coordinate { x, y: y+1 }); }
            }
            // South and east
            'F' => {
                if x + 1 < width { self.add_edge(Coordinate { x, y }, Coordinate { x: x+1, y }); }
                if y + 1 > height { self.add_edge(Coordinate { x, y }, Coordinate { x, y: y+1 }); }
            }
            '.' => {  },
            'S' => { self.update_start(Coordinate { x, y }); },
            _ => { println!("Unknown character: {}", char) }
        }
    }

    fn new(coordinates: Vec<Coordinate>, start: Coordinate) -> Self {
        let vertices = coordinates.clone();
        let vertices_count = vertices.len();
        Graph {
            start,
            vertices,
            adjacency_list: vec![Vec::new(); vertices_count],
        }
    }

    fn update_start(&mut self, start: Coordinate) {
        self.start = start;
    }

    fn add_edge(&mut self, coord1: Coordinate, coord2: Coordinate) {
        let index1 = self.vertices.iter().position(|&c| c == coord1).unwrap();
        let index2 = self.vertices.iter().position(|&c| c == coord2).unwrap();

        self.adjacency_list[index1].push(index2);
        self.adjacency_list[index2].push(index1); // For undirected graph
    }

    fn remove_edge(&mut self, coord1: Coordinate, coord2: Coordinate) {
        let index1 = self.vertices.iter().position(|&c| c == coord1).unwrap();
        let index2 = self.vertices.iter().position(|&c| c == coord2).unwrap();

        self.adjacency_list[index1].retain(|&x| x != index2);
        self.adjacency_list[index2].retain(|&x| x != index1); // For undirected graph
    }

    fn bfs(&self, start: Coordinate) -> Option<Vec<Coordinate>> {
        let start_index = self.vertices.iter().position(|&c| c == start)?;
        let end_index = self.vertices.iter().position(|&c| c == self.start)?;

        let mut visited = HashSet::new();
        let mut queue = VecDeque::new();
        let mut parent: HashMap<usize, usize> = HashMap::new();

        visited.insert(start_index);
        queue.push_back(start_index);

        while !queue.is_empty() {
            let current_vertex = queue.pop_front().unwrap();

            if current_vertex == end_index {
                // Reconstruct path if the end vertex is reached
                return Some(self.reconstruct_path(start_index, end_index, &parent));
            }

            for &neighbor in &self.adjacency_list[current_vertex] {
                if !visited.contains(&neighbor) {
                    visited.insert(neighbor);
                    parent.insert(neighbor, current_vertex);
                    queue.push_back(neighbor);
                }
            }
        }

        None // No path found
    }

    fn reconstruct_path(&self, start: usize, end: usize, parent: &HashMap<usize, usize>) -> Vec<Coordinate> {
        let mut path = Vec::new();
        let mut current_vertex = end;

        while current_vertex != start {
            path.push(self.vertices[current_vertex]);
            current_vertex = *parent.get(&current_vertex).unwrap();
        }

        path.push(self.vertices[start]);
        path.reverse();

        path
    }
}

fn main() {
    // let mut graph = Graph::from_file("input.txt");
    //
    // let mut start = graph.start;
    //
    // start.x -= 1;
    // graph.remove_edge(start, graph.start);
    // if let Some(path) = graph.bfs(start) {
    //     println!("Length: {} : Shortest: {}", path.len(), path.len()/2);
    // } else {
    //     println!("No path found from {:?} to {:?}", graph.start, start);
    // }
    part_1();
    part_2();
}

#[cfg(test)]
mod tests {
    use crate::Graph;

    #[test]
    fn part_1_test() {
        let mut graph = Graph::from_file("test.txt");
        let mut graph2 = Graph::from_file("test2.txt");

        let mut start = graph.start;
        start.x += 1;
        let mut start2 = graph2.start;
        start2.x += 1;

        graph.remove_edge(start, graph.start);
        graph2.remove_edge(start2, graph2.start);

        if let Some(path) = graph.bfs(start) {
            println!("Shortest: {}", path.len()/2);
        } else {
            println!("No path found from {:?} to {:?}", graph.start, start);
        }

        if let Some(path) = graph2.bfs(start2) {
            println!("Shortest: {}", path.len()/2);
        } else {
            println!("No path found from {:?} to {:?}", graph2.start, start2);
        }

        start2.x -= 1;
        start2.y += 1;
        graph2.remove_edge(start2, graph2.start);
        println!("Start: {:?}", start2);
        if let Some(path) = graph2.bfs(start2) {
            println!("Shortest: {}", path.len()/2);
        } else {
            println!("No path found from {:?} to {:?}", graph2.start, start2);
        }
    }
}