package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	var f_list []int
	var s_list []int

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

		f_list = append(f_list, a)
		s_list = append(s_list, b)
	}

	sort.Ints(f_list)
	sort.Ints(s_list)

	for ind, value := range f_list {
		dist := value - s_list[ind]
		if dist < 0 {
			dist = -dist
		}
		sum += dist
	}

	fmt.Println("The result is: %d", sum)
}
