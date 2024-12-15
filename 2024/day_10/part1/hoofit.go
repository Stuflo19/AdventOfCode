package part1

import (
	"fmt"
	"image"

	"stuflo19/aoc2024/helpers"
)

type Path struct {
	coordinate image.Point
	height     int
}

func newPath(coordinate image.Point, height int) Path {
	return Path{
		coordinate: coordinate,
		height:     height,
	}
}

type Bounds struct {
	minX int
	maxX int
	minY int
	maxY int
}

func buildGraph(contents [][]int) (map[string][]Path, []Path) {
	startingLocations := []Path{}
	graph := map[string][]Path{}
	bounds := Bounds{
		minX: 0,
		maxX: len(contents[0]) - 1,
		minY: 0,
		maxY: len(contents) - 1,
	}

	for y, line := range contents {
		for x, height := range line {
			coordinate := image.Pt(x, y)
			if height == 0 {
				startingLocations = append(startingLocations, newPath(coordinate, height))
			} else if height == 9 {
				continue
			}

			if x > bounds.minX {
				nextHeight := contents[y][x-1]
				if nextHeight == height+1 {
					graph[coordinate.String()] = append(
						graph[coordinate.String()],
						newPath(image.Pt(x-1, y), nextHeight),
					)
				}
			}
			if x < bounds.maxX {
				nextHeight := contents[y][x+1]
				if nextHeight == height+1 {
					graph[coordinate.String()] = append(
						graph[coordinate.String()],
						newPath(image.Pt(x+1, y), nextHeight),
					)
				}
			}
			if y > bounds.minY {
				nextHeight := contents[y-1][x]
				if nextHeight == height+1 {
					graph[coordinate.String()] = append(
						graph[coordinate.String()],
						newPath(image.Pt(x, y-1), nextHeight),
					)
				}
			}
			if y < bounds.maxY {
				nextHeight := contents[y+1][x]
				if nextHeight == height+1 {
					graph[coordinate.String()] = append(
						graph[coordinate.String()],
						newPath(image.Pt(x, y+1), nextHeight),
					)
				}
			}
		}
	}

	return graph, startingLocations
}

func BFS(graph map[string][]Path, start Path) int {
	score := 0
	visited := make(map[string]bool)
	queue := []Path{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.coordinate.String()] {
			continue
		}

		visited[current.coordinate.String()] = true
		if current.height == 9 {
			score += 1
		}

		for _, neighbor := range graph[current.coordinate.String()] {
			if !visited[neighbor.coordinate.String()] {
				queue = append(queue, neighbor)
			}
		}
	}

	return score
}

func Part1(filename string) {
	fileContents := helpers.ReadFileWithSeparatorAsInt(filename, "")

	graph, startLocations := buildGraph(fileContents)

	total := 0

	for _, start := range startLocations {
		total += BFS(graph, start)
	}

	fmt.Println("Part 1:", total)
}
