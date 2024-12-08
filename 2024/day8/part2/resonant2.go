package part2

import (
	"fmt"

	"stuflo19/aoc2024/helpers"
)

type Position struct {
	X int
	Y int
}

func (p Position) toString() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type Bounds struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func getLocationIncreasing(
	position Position,
	p Position,
	bounds Bounds,
) (Position, error) {
	newX := position.X*2 - p.X
	newY := position.Y*2 - p.Y

	newPosition := Position{X: int(newX), Y: int(newY)}

	if newPosition.X >= bounds.MaxX || newPosition.X < bounds.MinX ||
		newPosition.Y >= bounds.MaxY ||
		newPosition.Y < bounds.MinY {
		return Position{X: 0, Y: 0}, fmt.Errorf("Out of bounds")
	}

	return newPosition, nil
}

func getLocationDecreasing(
	position Position,
	p Position,
	bounds Bounds,
) (Position, error) {
	newX := p.X*2 - position.X
	newY := p.Y*2 - position.Y

	newPosition := Position{X: int(newX), Y: int(newY)}

	if newPosition.X >= bounds.MaxX ||
		newPosition.X < bounds.MinX ||
		newPosition.Y >= bounds.MaxY ||
		newPosition.Y < bounds.MinY {
		return Position{X: 0, Y: 0}, fmt.Errorf("Out of bounds")
	}

	return newPosition, nil
}

func getAllLocationsIncreasing(
	position Position,
	p Position,
	bounds Bounds,
) []Position {
	foundAntinodes := make([]Position, 0, 50)
	antinode, err := getLocationIncreasing(position, p, bounds)

	for err == nil {
		foundAntinodes = append(foundAntinodes, antinode)
		p = position
		position = antinode
		antinode, err = getLocationIncreasing(position, p, bounds)
	}

	return foundAntinodes
}

func getAllLocationsDecreasing(
	position Position,
	p Position,
	bounds Bounds,
) []Position {
	foundAntinodes := make([]Position, 0, 50)
	antinode, err := getLocationDecreasing(position, p, bounds)

	for err == nil {
		foundAntinodes = append(foundAntinodes, antinode)
		position = p
		p = antinode
		antinode, err = getLocationDecreasing(position, p, bounds)
	}

	return foundAntinodes
}

func calculateAntinodeLocationHorizontal(
	p Position,
	position Position,
	bounds Bounds,
) []Position {
	foundAntinodes := make([]Position, 0, 50)

	decreasingAntinodes := getAllLocationsDecreasing(position, p, bounds)

	for _, antinode := range decreasingAntinodes {
		foundAntinodes = append(foundAntinodes, antinode)
	}

	increasingAntinodes := getAllLocationsIncreasing(position, p, bounds)

	for _, antinode := range increasingAntinodes {
		foundAntinodes = append(foundAntinodes, antinode)
	}

	return foundAntinodes
}

func Part2(filename string) {
	fileContents := helpers.ReadFileWithSeparator(filename, "")
	bounds := Bounds{MinX: 0, MaxX: len(fileContents[0]), MinY: 0, MaxY: len(fileContents)}
	antennas := make(map[string][]Position)
	antinodes := make(map[string]Position)

	for col, line := range fileContents {
		for row, char := range line {
			if char != "." {
				newPosition := Position{X: row, Y: col}
				antennas[char] = append(antennas[char], newPosition)
				antinodes[newPosition.toString()] = newPosition
			}
		}
	}

	for _, positions := range antennas {
		for _, position := range positions {
			for _, p := range positions {
				if p.X == position.X && p.Y == position.Y {
					continue
				}

				foundAntinodes := calculateAntinodeLocationHorizontal(
					position,
					p,
					bounds,
				)

				for _, antinode := range foundAntinodes {
					antinodes[antinode.toString()] = antinode
				}
			}
		}
	}
	total := len(antinodes)

	fmt.Println("Part 2:", total)
}
