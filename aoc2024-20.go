package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var matrix [][]int
	var rowCount int
	var trailheads [][2]int
	sum := 0
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		line := reader.Text()
		if line == "" {
			break
		}

		var row []int
		for _, char := range line {
			row = append(row, int(char-'0'))
		}
		matrix = append(matrix, row)
		rowCount++
	}

	if err := reader.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
	}

	for _, trailhead := range trailheads {
		sum += bfs(matrix, trailhead, 0)
	}

	fmt.Println("Sum of paths to target 9:", sum)
}

func bfs(matrix [][]int, start [2]int, target int) int {
	max := len(matrix)

	if start[0] < 0 || start[0] >= max || start[1] < 0 || start[1] >= max {
		return 0
	}

	num := matrix[start[0]][start[1]]
	if num != target {
		return 0
	}

	if target == 9 {
		return 1
	}

	return bfs(matrix, [2]int{start[0] - 1, start[1]}, target+1) +
		bfs(matrix, [2]int{start[0] + 1, start[1]}, target+1) +
		bfs(matrix, [2]int{start[0], start[1] - 1}, target+1) +
		bfs(matrix, [2]int{start[0], start[1] + 1}, target+1)
}
