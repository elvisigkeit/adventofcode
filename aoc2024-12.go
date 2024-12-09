package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	X      int
	Y      int
	marker rune
	next   *Direction
}

var (
	RIGHT = Direction{0, 1, '>', nil}
	UP    = Direction{-1, 0, '^', nil}
	DOWN  = Direction{1, 0, 'v', nil}
	LEFT  = Direction{0, -1, '<', nil}
)

func debugMatrix(matrix [][]rune) {
	fmt.Print("  ")
	for ind, _ := range matrix {
		fmt.Print(ind, " ")
	}
	fmt.Println()
	for indR, row := range matrix {
		fmt.Print(indR, " ")
		for _, cell := range row {
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func outOfBounds(x, y, len int) bool {
	return x < 0 || x >= len || y < 0 || y >= len
}

func wouldCloseCycle(matrix [][]rune, x, y int, dir Direction) bool {
	len := len(matrix)
	nextDir := *dir.next
	cycleMarker := (*nextDir.next).marker
	nextX, nextY := x+nextDir.X, y+nextDir.Y
	scanCycle := make(map[rune]map[Direction]struct{})
	// fmt.Printf("Scanning for cycle: %d %d\n", x, y)
	for !outOfBounds(nextX, nextY, len) {
		nextMarker := matrix[nextX][nextY]
		if nextMarker == '#' {
			nextX, nextY = nextX-nextDir.X, nextY-nextDir.Y
			nextDir = *nextDir.next
			cycleMarker = (*nextDir.next).marker
			nextX, nextY = nextX+nextDir.X, nextY+nextDir.Y
		}
		// fmt.Println("checked", nextX, nextY, "marker", nextMarker, "cycle marker", string(cycleMarker))

		scanDir := Direction{nextX, nextY, cycleMarker, nil}
		if scanCycle[cycleMarker] == nil {
			scanCycle[cycleMarker] = make(map[Direction]struct{})
		}
		if _, exists := scanCycle[cycleMarker][scanDir]; !exists {
			scanCycle[cycleMarker][scanDir] = struct{}{}
		} else {
			return true
		}

		if nextMarker == cycleMarker {
			return true
		}
		nextX, nextY = nextX+nextDir.X, nextY+nextDir.Y
	}
	return false
}

func sumPathToExit(matrix [][]rune, x, y, sum int, dir Direction) int {
	len := len(matrix)

	// debugMatrix(matrix)
	nextX, nextY := x+dir.X, y+dir.Y

	if outOfBounds(nextX, nextY, len) {
		return sum
	}

	if matrix[nextX][nextY] != '#' &&
		matrix[nextX][nextY] != 'X' &&
		wouldCloseCycle(matrix, x, y, dir) {
		sum += 1
	}

	for matrix[nextX][nextY] == '#' {
		dir = *dir.next
		nextX, nextY = x+dir.X, y+dir.Y
		matrix[x][y] = dir.marker
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
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		matrix = append(matrix, []rune(line))
	}

	for i := 0; i < len(matrix) && !found; i++ {
		for j := 0; j < len(matrix[i]) && !found; j++ {
			if matrix[i][j] == '^' {
				matrix[i][j] = 'X'
				sum = sumPathToExit(matrix, i, j, sum, UP)
				found = true
			}
		}
	}

	fmt.Println("Result: ", sum)
}
