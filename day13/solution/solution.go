package solution

import (
	"strconv"
	"strings"

	"github.com/danielstepan/adventofcode/internal/input"
)

type Machine struct {
	ax, ay, bx, by, px, py int
}

func NewMachine(ax, ay, bx, by, px, py int) *Machine {
	return &Machine{
		ax: ax,
		ay: ay,
		bx: bx,
		by: by,
		px: px,
		py: py,
	}
}

func calculateButtonPress(machine *Machine) (int, int) {
	det := (machine.ax*machine.by - machine.ay*machine.bx)

	A := (machine.px*machine.by - machine.py*machine.bx) / det
	B := (machine.ax*machine.py - machine.ay*machine.px) / det

	if (machine.ax*A+machine.bx*B) == machine.px && (machine.ay*A+machine.by*B) == machine.py {
		return A, B
	} else {
		return -1, -1
	}
}

func parseBlockToMachine(lines []string, priceOffset int) *Machine {
	var ax, ay, bx, by, px, py int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Button A:") {
			ax, ay = parseCoordinates(line)
		} else if strings.HasPrefix(line, "Button B:") {
			bx, by = parseCoordinates(line)
		} else if strings.HasPrefix(line, "Prize:") {
			px, py = parseCoordinates(line)
		}
	}

	return NewMachine(ax, ay, bx, by, px+priceOffset, py+priceOffset)
}

func parseCoordinates(line string) (int, int) {
	parts := strings.Split(line, ":")[1]
	coords := strings.Split(parts, ",")

	x := parseCoordinate(coords[0])
	y := parseCoordinate(coords[1])

	return x, y
}

func parseCoordinate(part string) int {
	part = strings.TrimSpace(part)
	if strings.Contains(part, "+") || strings.Contains(part, "=") {
		val := strings.Split(part, "+")
		if len(val) < 2 {
			val = strings.Split(part, "=")
		}
		num, _ := strconv.Atoi(strings.TrimSpace(val[1]))
		return num
	}
	return 0
}

func run(inputFilePath string, priceOffset int) (string, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return "", err
	}

	count := 0

	for i := 0; i < len(content); i = i + 4 {
		machine := parseBlockToMachine(content[i:i+3], priceOffset)
		A, B := calculateButtonPress(machine)
		if A == -1 && B == -1 {
			continue
		}

		count += A*3 + B
	}

	return strconv.Itoa(count), nil
}

func Part1(inputFilePath string) (string, error) {
	return run(inputFilePath, 0)
}

func Part2(inputFilePath string) (string, error) {
	return run(inputFilePath, 10000000000000)
}
