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
	flag := []rune("XMAS")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		matrix = append(matrix, []rune(line))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == flag[0] {
				sum += checkFlag(matrix, i, j, 1, 0, flag, 1)
				sum += checkFlag(matrix, i, j, 0, 1, flag, 1)
				sum += checkFlag(matrix, i, j, 1, 1, flag, 1)
				sum += checkFlag(matrix, i, j, 0, -1, flag, 1)
				sum += checkFlag(matrix, i, j, -1, 0, flag, 1)
				sum += checkFlag(matrix, i, j, -1, 1, flag, 1)
				sum += checkFlag(matrix, i, j, 1, -1, flag, 1)
				sum += checkFlag(matrix, i, j, -1, -1, flag, 1)
			}
		}
	}

	fmt.Println("Result: ", sum)
}

func checkFlag(matrix [][]rune, x, y, dx, dy int, flag []rune, index int) int {
	if index == len(flag) {
		return 1
	}

	x = x + dx
	y = y + dy

	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[x]) {
		return 0
	}

	if matrix[x][y] == flag[index] {
		return checkFlag(matrix, x, y, dx, dy, flag, index+1)
	}

	return 0
}
