package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()

		if line == "" {
			break
		}

		re := regexp.MustCompile(`mul\(\d+,\d+\)`)

		matches := re.FindAllString(line, -1)

		for _, match := range matches {
			nums := match[4 : len(match)-1]
			parts := strings.Split(nums, ",")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			sum += a * b
		}
	}

	fmt.Println("The result is: ", sum)
}
