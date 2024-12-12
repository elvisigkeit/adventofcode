package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Subtract(other Point) Point {
	return Point{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	matrix := [][]rune{}
	antinodes := make(map[Point]struct{})

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		row := []rune(input)
		matrix = append(matrix, row)
	}

	runeMap := make(map[rune][]Point)

	for i, row := range matrix {
		for j, r := range row {
			if r != '.' {
				runeMap[r] = append(runeMap[r], Point{X: j, Y: i})
			}
		}
	}

	for _, points := range runeMap {
		for i := 0; i < len(points); i++ {
			for j := 0; j < len(points); j++ {
				if i == j {
					continue
				}

				dP := points[i].Subtract(points[j])
				antinode := points[i]
				for !(antinode.X < 0 || antinode.X >= len(matrix) || antinode.Y < 0 || antinode.Y >= len(matrix[0])) {
					antinodes[antinode] = struct{}{}
					antinode = antinode.Add(dP)
				}
			}
		}
	}

	fmt.Println("Result: ", len(antinodes))
}
