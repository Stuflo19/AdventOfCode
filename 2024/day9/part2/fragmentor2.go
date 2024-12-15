package part2

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

			for i := 0; i < asInt; i++ {
				lineDetails = append(lineDetails, types.NewBlock(-1))
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

func defrag(lineDetails []types.Block) []types.Block {
	currId := lineDetails[len(lineDetails)-1].GetID()

NextId:
	for currId > 0 {
		indexesOfCurrId := []int{}
		indexesOfCurrFreeSpace := []int{}

		for i, block := range lineDetails {
			if block.GetID() == currId {
				indexesOfCurrId = append(indexesOfCurrId, i)
			} else if block.GetID() != currId && len(indexesOfCurrId) > 0 {
				break
			}
		}

		for i, block := range lineDetails {
			if block.GetID() == -1 {
				indexesOfCurrFreeSpace = append(indexesOfCurrFreeSpace, i)
			} else if i > indexesOfCurrId[0] {
				currId--
				continue NextId
			} else {
				indexesOfCurrFreeSpace = []int{}
			}

			if len(indexesOfCurrId) > 0 && len(indexesOfCurrFreeSpace) == len(indexesOfCurrId) {
				// swap
				for i := 0; i < len(indexesOfCurrId); i++ {
					temp := lineDetails[indexesOfCurrId[i]]
					lineDetails[indexesOfCurrId[i]] = lineDetails[indexesOfCurrFreeSpace[i]]
					lineDetails[indexesOfCurrFreeSpace[i]] = temp
				}
				currId--
				continue NextId
			}
		}
		currId--
	}

	return lineDetails
}

func Part2(filename string) {
	fileContents := helpers.ReadFileWithSeparator(filename, "")

	lineDetails, _ := generateLineDetails(fileContents[0])

	lineDetails = defrag(lineDetails)

	total := 0

	for idx, block := range lineDetails {
		if block.GetID() == -1 {
			continue
		}
		total += block.GetID() * idx
	}

	fmt.Println("Part 2:", total)
}
