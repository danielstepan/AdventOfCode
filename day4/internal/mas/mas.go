package mas

func FoundMAS(letters [][]string, i, j int) bool {
	return searchDiagonalLeftToRightXMAS(letters, i, j) && searchDiagonalRightToLeftXMAS(letters, i, j)
}

func searchDiagonalLeftToRightXMAS(letters [][]string, i, j int) bool {
	if (i-1 < 0 || j-1 < 0) || (i+1 >= len(letters) || j+1 >= len(letters[i])) {
		return false
	}

	upperLeft := letters[i-1][j-1]
	lowerRight := letters[i+1][j+1]

	return (upperLeft == "M" && lowerRight == "S") || (upperLeft == "S" && lowerRight == "M")
}

func searchDiagonalRightToLeftXMAS(letters [][]string, i, j int) bool {
	if (i-1 < 0 || j+1 >= len(letters[i])) || (i+1 >= len(letters) || j-1 < 0) {
		return false
	}

	upperRight := letters[i-1][j+1]
	lowerLeft := letters[i+1][j-1]

	return (upperRight == "M" && lowerLeft == "S") || (upperRight == "S" && lowerLeft == "M")
}
