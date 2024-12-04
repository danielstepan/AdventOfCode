package xmas

var searchFunctions = []func([][]string, int, int) bool{
	searchRightHorizontalXMAS,
	searchLeftHorizontalXMAS,
	searchVerticalUpXMAS,
	searchVerticalDownXMAS,
	searchDiagonalUpRightXMAS,
	searchDiagonalUpLeftXMAS,
	searchDiagonalDownRightXMAS,
	searchDiagonalDownLeftXMAS,
}

func FoundXMASCount(letters [][]string, i int, j int) int {
	count := 0
	for _, searchFunction := range searchFunctions {
		if searchFunction(letters, i, j) {
			count++
		}
	}
	return count
}

func searchRightHorizontalXMAS(letters [][]string, i int, j int) bool {
	if j+3 >= len(letters[i]) {
		return false
	}

	if letters[i][j+1] == "M" && letters[i][j+2] == "A" && letters[i][j+3] == "S" {
		return true
	}
	return false
}

func searchLeftHorizontalXMAS(letters [][]string, i int, j int) bool {
	if j-3 < 0 {
		return false
	}

	if letters[i][j-1] == "M" && letters[i][j-2] == "A" && letters[i][j-3] == "S" {
		return true
	}
	return false
}

func searchVerticalUpXMAS(letters [][]string, i int, j int) bool {
	if i-3 < 0 {
		return false
	}

	if letters[i-1][j] == "M" && letters[i-2][j] == "A" && letters[i-3][j] == "S" {
		return true
	}
	return false
}

func searchVerticalDownXMAS(letters [][]string, i int, j int) bool {
	if i+3 >= len(letters) {
		return false
	}

	if letters[i+1][j] == "M" && letters[i+2][j] == "A" && letters[i+3][j] == "S" {
		return true
	}
	return false
}

func searchDiagonalUpRightXMAS(letters [][]string, i int, j int) bool {
	if i-3 < 0 || j+3 >= len(letters[i]) {
		return false
	}

	if letters[i-1][j+1] == "M" && letters[i-2][j+2] == "A" && letters[i-3][j+3] == "S" {
		return true
	}
	return false
}

func searchDiagonalUpLeftXMAS(letters [][]string, i int, j int) bool {
	if i-3 < 0 || j-3 < 0 {
		return false
	}

	if letters[i-1][j-1] == "M" && letters[i-2][j-2] == "A" && letters[i-3][j-3] == "S" {
		return true
	}
	return false
}

func searchDiagonalDownRightXMAS(letters [][]string, i int, j int) bool {
	if i+3 >= len(letters) || j+3 >= len(letters[i]) {
		return false
	}

	if letters[i+1][j+1] == "M" && letters[i+2][j+2] == "A" && letters[i+3][j+3] == "S" {
		return true
	}
	return false
}

func searchDiagonalDownLeftXMAS(letters [][]string, i int, j int) bool {
	if i+3 >= len(letters) || j-3 < 0 {
		return false
	}

	if letters[i+1][j-1] == "M" && letters[i+2][j-2] == "A" && letters[i+3][j-3] == "S" {
		return true
	}
	return false
}
