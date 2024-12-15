package part2

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"

	day14 "stuflo19/aoc2024/day_14"
	"stuflo19/aoc2024/helpers"
)

func createGrid(maxX, maxY int) [][]string {
	grid := [][]string{}

	for y := 0; y < maxY; y++ {
		row := make([]string, maxX)
		for x := 0; x < maxX; x++ {
			row[x] = "."
		}
		grid = append(grid, row)
	}

	return grid
}

func writeGridToPng(grid [][]string, iteration int) {
	height := len(grid)
	width := len(grid[0])

	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == "." {
				grayValue := uint8(255)
				img.SetGray(x, y, color.Gray{Y: grayValue})
			} else {
				grayValue := uint8(0)
				img.SetGray(x, y, color.Gray{Y: grayValue})
			}
		}
	}

	file, err := os.Create(
		fmt.Sprintf(
			"/Users/stu/Development/AdventOfCode/2024/day_14/part2/iterations/grid_%d.png",
			iteration,
		),
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func Part2(filename string, iterations int, maxX int, maxY int) {
	contents := helpers.ReadFile(filename)
	robots := []day14.Robot{}

	grid := createGrid(maxX, maxY)

	for _, line := range contents {
		robots = append(robots, day14.NewRobot(line))
	}

	for i := 0; i < iterations; i++ {
		for idx := range robots {
			currValue := grid[robots[idx].GetPos().Y][robots[idx].GetPos().X]
			if currValue == "1" || currValue == "." {
				grid[robots[idx].GetPos().Y][robots[idx].GetPos().X] = "."
			} else {
				intCurr, _ := strconv.Atoi(currValue)
				grid[robots[idx].GetPos().Y][robots[idx].GetPos().X] = strconv.Itoa(intCurr - 1)
			}

			robots[idx].MoveTimes(1, maxX, maxY)

			currValue = grid[robots[idx].GetPos().Y][robots[idx].GetPos().X]
			if currValue == "." {
				grid[robots[idx].GetPos().Y][robots[idx].GetPos().X] = "1"
			} else {
				intCurr, _ := strconv.Atoi(currValue)
				grid[robots[idx].GetPos().Y][robots[idx].GetPos().X] = strconv.Itoa(intCurr + 1)
			}
		}

		writeGridToPng(grid, i)
	}
}
