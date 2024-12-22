package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	Count int
	Next  []int
}

func countDigitsLog(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	return int(math.Floor(math.Log10(float64(n))) + 1)
}

func calculateNext(value int, numDigits int) []int {
	if value == 0 {
		return []int{1}
	}
	if numDigits%2 == 0 {
		digits := strconv.Itoa(value)
		leftHalf := digits[:len(digits)/2]
		rightHalf := digits[len(digits)/2:]
		newLeft, _ := strconv.Atoi(leftHalf)
		newRight, _ := strconv.Atoi(rightHalf)
		return []int{newLeft, newRight}
	}
	return []int{value * 2024}
}

func addStone(stones map[int]Stone, value, count int) {
	if stone, exists := stones[value]; !exists {
		numDigits := countDigitsLog(value)
		stones[value] = Stone{Count: count, Next: calculateNext(value, numDigits)}
	} else {
		stone.Count += count
		stones[value] = stone
	}
}

func countStones(stones map[int]Stone) int {
	sum := 0
	for _, stone := range stones {
		sum += stone.Count
	}
	return sum
}

func main() {
	steps := 75
	stones := map[int]Stone{}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numbers := strings.Split(input, " ")
	for _, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		addStone(stones, num, 1)
	}

	for i := 0; i < steps; i++ {
		newStones := map[int]Stone{}
		for _, stone := range stones {
			for _, nextValue := range stone.Next {
				addStone(newStones, nextValue, stone.Count)
			}
		}

		stones = newStones
		fmt.Println("Step: ", i)
		// for stone, _ := range stones {
		// 	fmt.Print(stone, " ")
		// }
		// fmt.Println()
		fmt.Println("Count: ", countStones(stones))
	}

	fmt.Println("\nCount: ", countStones(stones))
}
