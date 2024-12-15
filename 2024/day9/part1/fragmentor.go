package part1

import (
	"fmt"
	"strconv"

	"stuflo19/aoc2024/day9/types"
	"stuflo19/aoc2024/helpers"
)

func generateLineDetails(line []string) ([]types.Block, []types.FreeSpace) {
	lineDetails := []types.Block{}
	freeSpace := []types.FreeSpace{}
	id := 0
	currPos := 0

	for idx, char := range line {
		asInt, _ := strconv.Atoi(char)

		if (idx+1)%2 == 0 {
			// on even indexes we have a free space
			if asInt == 0 {
				continue
			}
			freeSpace = append(freeSpace, types.NewFreeSpace(currPos, (currPos-1)+asInt))
			currPos += asInt
		} else {
			// on odd indexes we have a block
			for i := 0; i < asInt; i++ {
				lineDetails = append(lineDetails, types.NewBlock(id))
			}

			currPos += asInt
			id++
		}
	}

	return lineDetails, freeSpace
}

func defrag(lineDetails []types.Block, freeSpace []types.FreeSpace) []types.Block {
	for _, free := range freeSpace {
		for i := free.GetStartPos(); i <= free.GetEndPos(); i++ {
			if i >= len(lineDetails) {
				return lineDetails
			}

			lastItem := []types.Block{lineDetails[len(lineDetails)-1]}
			lineDetails = append(
				lineDetails[:i],
				append(lastItem, lineDetails[i:len(lineDetails)-1]...)...)
		}
	}

	return lineDetails
}

func Part1(filename string) {
	fileContents := helpers.ReadFileWithSeparator(filename, "")

	lineDetails, freeSpace := generateLineDetails(fileContents[0])

	lineDetails = defrag(lineDetails, freeSpace)

	total := 0

	for idx, block := range lineDetails {
		total += block.GetID() * idx
	}

	fmt.Println("Part 1:", total)
}
