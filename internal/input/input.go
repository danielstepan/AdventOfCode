package input

import (
	"bufio"
	"os"
)

func ReadFileLines(name string) (content []string, err error) {
	file, err := os.Open(name)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content, nil
}
