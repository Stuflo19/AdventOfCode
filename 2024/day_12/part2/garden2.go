package part2

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
	corners := 0
	crop := garden[plot[0].Y][plot[0].X]

	for _, point := range plot {
		north := image.Point{X: point.X, Y: point.Y - 1}
		south := image.Point{X: point.X, Y: point.Y + 1}
		west := image.Point{X: point.X - 1, Y: point.Y}
		east := image.Point{X: point.X + 1, Y: point.Y}

		northEdge := false
		southEdge := false
		westEdge := false
		eastEdge := false

		northEastEdge := false
		northWestEdge := false
		southEastEdge := false
		southWestEdge := false

		if north.Y < 0 || garden[north.Y][north.X] != crop {
			northEdge = true
		}
		if south.Y >= len(garden) || garden[south.Y][south.X] != crop {
			southEdge = true
		}
		if west.X < 0 || garden[west.Y][west.X] != crop {
			westEdge = true
		}
		if east.X >= len(garden[0]) || garden[east.Y][east.X] != crop {
			eastEdge = true
		}

		if north.Y < 0 || east.X >= len(garden[0]) || garden[north.Y][east.X] != crop {
			northEastEdge = !(north.Y < 0 || east.X >= len(garden[0]))
		}
		if north.Y < 0 || west.X < 0 || garden[north.Y][west.X] != crop {
			northWestEdge = !(north.Y < 0 || west.X < 0)
		}
		if south.Y >= len(garden) || east.X >= len(garden[0]) || garden[south.Y][east.X] != crop {
			southEastEdge = !(south.Y >= len(garden) || east.X >= len(garden[0]))
		}
		if south.Y >= len(garden) || west.X < 0 || garden[south.Y][west.X] != crop {
			southWestEdge = !(south.Y >= len(garden) || west.X < 0)
		}

		if northEdge && westEdge {
			corners += 1
		}

		if northEdge && eastEdge {
			corners += 1
		}

		if southEdge && westEdge {
			corners += 1
		}

		if southEdge && eastEdge {
			corners += 1
		}

		if !northEdge && !westEdge && northWestEdge {
			corners += 1
		}
		if !northEdge && !eastEdge && northEastEdge {
			corners += 1
		}
		if !southEdge && !westEdge && southWestEdge {
			corners += 1
		}
		if !southEdge && !eastEdge && southEastEdge {
			corners += 1
		}
	}

	return corners
}

func Part2(filename string) {
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

	fmt.Println("Part 2:", total)
}
