package part1

import (
	"fmt"
	"image"

	"stuflo19/aoc2024/helpers"
)

func dfs(
	garden [][]string,
	crop string,
	point image.Point,
	visited *map[string]bool,
) []image.Point {
	maxX := len(garden[0])
	maxY := len(garden)
	plot := []image.Point{}
	queue := []image.Point{point}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.X < 0 || current.X >= maxX || current.Y < 0 || current.Y >= maxY {
			continue
		} else if crop != garden[current.Y][current.X] {
			continue
		} else if (*visited)[current.String()] {
			continue
		}

		(*visited)[current.String()] = true
		plot = append(plot, current)
		queue = append(queue, image.Pt(current.X-1, current.Y))
		queue = append(queue, image.Pt(current.X+1, current.Y))
		queue = append(queue, image.Pt(current.X, current.Y-1))
		queue = append(queue, image.Pt(current.X, current.Y+1))
	}

	return plot
}

func findPlots(garden [][]string) [][]image.Point {
	visited := make(map[string]bool)
	plots := [][]image.Point{}

	for y, col := range garden {
		for x, crop := range col {
			point := image.Pt(x, y)
			if !visited[point.String()] {
				newPlot := dfs(garden, crop, point, &visited)
				plots = append(plots, newPlot)
			}
		}
	}

	return plots
}

func getPerimiter(garden [][]string, plot []image.Point) int {
	perimiter := 0
	crop := garden[plot[0].Y][plot[0].X]

	for _, point := range plot {
		up := image.Point{X: point.X, Y: point.Y - 1}
		down := image.Point{X: point.X, Y: point.Y + 1}
		left := image.Point{X: point.X - 1, Y: point.Y}
		right := image.Point{X: point.X + 1, Y: point.Y}

		if up.Y < 0 || garden[up.Y][up.X] != crop {
			perimiter += 1
		}
		if down.Y >= len(garden) || garden[down.Y][down.X] != crop {
			perimiter += 1
		}
		if left.X < 0 || garden[left.Y][left.X] != crop {
			perimiter += 1
		}
		if right.X >= len(garden[0]) || garden[right.Y][right.X] != crop {
			perimiter += 1
		}
	}

	return perimiter
}

func Part1(filename string) {
	garden := helpers.ReadFileWithSeparator(filename, "")

	plots := findPlots(garden)

	total := 0
	for _, plot := range plots {
		area := len(plot)
		perimeter := getPerimiter(garden, plot)
		fmt.Println("Current crop", garden[plot[0].Y][plot[0].X])
		fmt.Println("Area", area)
		fmt.Println("Perimeter", perimeter)
		total += area * perimeter
	}

	fmt.Println("Part 1:", total)
}
