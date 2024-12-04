package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var matrix [][]rune
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		matrix = append(matrix, []rune(line))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'A' {
				sum += checkXMAS(matrix, i, j)
			}
		}
	}

	fmt.Println("Result: ", sum)
}

func checkXMAS(matrix [][]rune, x, y int) int {
	if x-1 < 0 || x+1 >= len(matrix) || y-1 < 0 || y+1 >= len(matrix[x]) {
		return 0
	}

	sum := 0

	sum += checkMAS(matrix, x, y, 0)
	sum += checkMAS(matrix, x, y, 1)
	sum += checkMAS(matrix, x, y, 2)
	sum += checkMAS(matrix, x, y, 3)

	return sum
}

func checkMAS(matrix [][]rune, x, y, version int) int {
	posNum := 4
	pos := [][2]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	mas := []rune("MMSS")

	if len(pos) != posNum || len(mas) != posNum {
		panic("Invalid MAS positions length")
	}

	for i := 0; i < posNum; i++ {
		index := (i + version) % posNum
		nx, ny := x+pos[index][0], y+pos[index][1]
		if matrix[nx][ny] != mas[i] {
			return 0
		}
	}

	return 1
}
