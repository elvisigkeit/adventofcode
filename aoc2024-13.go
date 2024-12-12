package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	sum := 0

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var numbers []int
		if input == "" {
			break
		}

		parts := strings.Split(input, ":")
		firstNumber, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numberStrings := strings.Fields(parts[1])
		for _, numStr := range numberStrings {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}

		left := validAddition(numbers[0], firstNumber, numbers, 1, 'x')
		if left {
			sum += firstNumber
		} else {
			right := validAddition(numbers[0], firstNumber, numbers, 1, '+')
			if right {
				sum += firstNumber
			}
		}
	}
	fmt.Println("The sum of all valid numbers is:", sum)
}

func validAddition(acc, target int, numbers []int, pos int, sign rune) bool {
	if pos == len(numbers) {
		return false
	}

	result := 0
	number := numbers[pos]
	if sign == 'x' {
		result = acc * number
	} else if sign == '+' {
		result = acc + number
	} else {
		panic("validAddition needs one of these signs: 'x' or '+'")
	}

	if result == target {
		return true
	} else if result < target {
		valid := validAddition(result, target, numbers, pos+1, 'x')
		if !valid {
			return validAddition(result, target, numbers, pos+1, '+')
		} else {
			return true
		}
	}

	return false
}
