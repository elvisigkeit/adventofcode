package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	f_map := make(map[int]int)
	s_map := make(map[int]int)

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()

		if line == "" {
			fmt.Println("Calculating...")
			break
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter exactly two integers.")
			continue
		}

		a, err1 := strconv.Atoi(parts[0])
		b, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error: Both inputs must be integers.")
			break
		}

		if value, exists := f_map[a]; exists {
			f_map[a] = value + 1
		} else {
			f_map[a] = 1
		}

		if value, exists := s_map[b]; exists {
			s_map[b] = value + 1
		} else {
			s_map[b] = 1
		}
	}

	for num, count := range f_map {
		sum += count * num * s_map[num]
	}

	fmt.Println("The result is: %d", sum)
}
