package solution

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/danielstepan/adventofcode/internal/input"
)

func prepareInput(inputFilePath string) (string, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return "", fmt.Errorf("could not read file: %w", err)
	}

	return content[0], nil
}

func Part1(inputFilePath string) (string, error) {
	diskMap, err := prepareInput(inputFilePath)
	if err != nil {
		return "", fmt.Errorf("could not prepare input: %w", err)
	}

	fragmentedData := convertDiskMap(diskMap)
	defragmentedData := defragmentDataByBlocks(fragmentedData)
	checksum := calculateChecksum(defragmentedData)

	return strconv.Itoa(checksum), nil
}

func Part2(inputFilePath string) (string, error) {
	diskMap, err := prepareInput(inputFilePath)
	if err != nil {
		return "", fmt.Errorf("could not prepare input: %w", err)
	}

	fragmentedData := convertDiskMap(diskMap)
	defragmentedData := defragmentDataByFiles(fragmentedData)
	checksum := calculateChecksum(defragmentedData)

	return strconv.Itoa(checksum), nil
}

func defragmentDataByBlocks(fragmentedData []int) []int {
	defragmentedData := make([]int, len(fragmentedData))
	copy(defragmentedData, fragmentedData)

	for i := len(defragmentedData) - 1; i >= 0; i-- {
		if defragmentedData[i] != -1 {
			firstSpace := slices.Index(defragmentedData, -1)
			if firstSpace > i {
				break
			}
			defragmentedData[firstSpace], defragmentedData[i] = defragmentedData[i], defragmentedData[firstSpace]
		}
	}
	return defragmentedData
}

func defragmentDataByFiles(fragmentedData []int) []int {
	defragmentedData := make([]int, len(fragmentedData))
	copy(defragmentedData, fragmentedData)

	for i := len(defragmentedData) - 1; i >= 0; i-- {
		if defragmentedData[i] == -1 {
			continue
		}

		// detect data block sequence
		blockId := defragmentedData[i]
		blockCount := 0
		for j := i; j >= 0 && defragmentedData[j] == blockId; j-- {
			blockCount++
		}
		// start looking for empty block sequence
		for j := 0; j < len(defragmentedData); j++ {
			if defragmentedData[j] != -1 {
				continue
			}
			if j >= i {
				break
			}

			// detect empty block sequence
			emptyBlockCount := 0
			for k := j; k < len(defragmentedData) && defragmentedData[k] == -1; k++ {
				emptyBlockCount++
			}

			// move data block sequence
			if emptyBlockCount >= blockCount {
				for k := 0; k < blockCount; k++ {
					defragmentedData[j+k] = blockId
					defragmentedData[i-k] = -1
				}
				break
			}
			// move current index behind the found empty block
			j += emptyBlockCount - 1
		}
		// move current index behind the found data block
		i -= blockCount - 1
	}
	return defragmentedData
}

func convertDiskMap(diskMap string) []int {
	var fragmentedData []int
	var currentDataBlockId int

	for i, blockCount := range diskMap {
		blockCountInt, _ := strconv.Atoi(string(blockCount))
		if i%2 == 0 {
			data := currentDataBlockId
			for j := 0; j < blockCountInt; j++ {
				fragmentedData = append(fragmentedData, data)
			}
			currentDataBlockId++
		} else {
			for j := 0; j < blockCountInt; j++ {
				fragmentedData = append(fragmentedData, -1)
			}
		}
	}
	return fragmentedData
}

func calculateChecksum(fragmentedData []int) int {
	var totalChecksum int
	for i := 0; i < len(fragmentedData); i++ {
		if fragmentedData[i] == -1 {
			continue
		}
		blockId := fragmentedData[i]
		totalChecksum += i * blockId
	}
	return totalChecksum
}
