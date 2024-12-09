package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	X    int
	Y    int
	next *Direction
}

var (
	RIGHT = Direction{0, 1, nil}
	UP    = Direction{-1, 0, nil}
	DOWN  = Direction{1, 0, nil}
	LEFT  = Direction{0, -1, nil}
)

func debugMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
	fmt.Println()
}

func outOfBounds(x, y, len int) bool {
	return x < 0 || x >= len || y < 0 || y >= len
}

func sumPathToExit(matrix [][]rune, x, y, sum int, dir Direction) int {
	len := len(matrix)

	if matrix[x][y] != 'X' {
		sum += 1
		matrix[x][y] = 'X'
	}

	// debugMatrix(matrix)
	nextX, nextY := x+dir.X, y+dir.Y

	if outOfBounds(nextX, nextY, len) {
		return sum
	}
	for matrix[nextX][nextY] == '#' {
		dir = *dir.next
		nextX, nextY = x+dir.X, y+dir.Y
		if outOfBounds(nextX, nextY, len) {
			return sum
		}
	}

	return sumPathToExit(matrix, nextX, nextY, sum, dir)
}

func main() {
	UP.next = &RIGHT
	RIGHT.next = &DOWN
	DOWN.next = &LEFT
	LEFT.next = &UP

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
			if matrix[i][j] == '^' {
				sum = sumPathToExit(matrix, i, j, sum, UP)
			}
		}
	}

	fmt.Println("Result: ", sum)
}
