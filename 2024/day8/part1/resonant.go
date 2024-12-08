package part1

import (
	"fmt"

	"stuflo19/aoc2024/helpers"
)

type Position struct {
	X int
	Y int
}

type Bounds struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func calculateNewAntinodeLocationUp(
	p Position,
	position Position,
	bounds Bounds,
) (Position, error) {
	newX := position.X - (p.X - position.X)
	newY := position.Y - (p.Y - position.Y)

	newPosition := Position{X: int(newX), Y: int(newY)}

	if newPosition.X >= bounds.MaxX ||
		newPosition.X < bounds.MinX ||
		newPosition.Y >= bounds.MaxY ||
		newPosition.Y < bounds.MinY {
		return Position{X: 0, Y: 0}, fmt.Errorf("Out of bounds")
	}

	return newPosition, nil
}

func calculateNewAntinodeLocationDown(
	p Position,
	position Position,
	bounds Bounds,
) (Position, error) {
	newX := position.X + (position.X - p.X)
	newY := position.Y + (position.Y - p.Y)

	newPosition := Position{X: int(newX), Y: int(newY)}

	if newPosition.X >= bounds.MaxX ||
		newPosition.X < bounds.MinX ||
		newPosition.Y >= bounds.MaxY ||
		newPosition.Y < bounds.MinY {
		return Position{X: 0, Y: 0}, fmt.Errorf("Out of bounds")
	}

	return newPosition, nil
}

func Part1(filename string) {
	fileContents := helpers.ReadFileWithSeparator(filename, "")
	copyOfFileContents := make([][]string, len(fileContents))
	bounds := Bounds{MinX: 0, MaxX: len(fileContents[0]), MinY: 0, MaxY: len(fileContents)}
	antennas := make(map[string][]Position)
	antinodes := make(map[string][]Position)

	for col, line := range fileContents {
		copyOfFileContents[col] = make([]string, len(line))
		copy(copyOfFileContents[col], line)
		for row, char := range line {
			if char != "." {
				antennas[char] = append(antennas[char], Position{X: row, Y: col})
			}
		}
	}

	for char, positions := range antennas {
		for _, position := range positions {

			for _, p := range positions {
				if p.X == position.X && p.Y == position.Y {
					continue
				}

				antinode, err := calculateNewAntinodeLocationUp(
					position,
					p,
					bounds,
				)

				if err != nil {
					continue
				}

				if copyOfFileContents[antinode.Y][antinode.X] != "#" {
					antinodes[char] = append(antinodes[char], antinode)
				}
				copyOfFileContents[antinode.Y][antinode.X] = "#"

				antinode, err = calculateNewAntinodeLocationDown(
					position,
					p,
					bounds,
				)

				if err != nil {
					continue
				}

				if copyOfFileContents[antinode.Y][antinode.X] != "#" {
					antinodes[char] = append(antinodes[char], antinode)
				}
				copyOfFileContents[antinode.Y][antinode.X] = "#"
			}
		}
	}

	fmt.Println("Bounds:", bounds)
	fmt.Println("Antennas:", antennas)
	fmt.Println("Antinodes:", antinodes)
	total := 0
	for _, v := range antinodes {
		total += len(v)
	}

	fmt.Println("Part 1:", total)

	for _, line := range copyOfFileContents {
		fmt.Println(line)
	}
}
