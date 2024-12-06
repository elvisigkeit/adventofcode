package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func existsInSet(set1, set2 map[string]struct{}) bool {
	for key := range set1 {
		if _, exists := set2[key]; exists {
			return true
		}
	}
	return false
}

func main() {
	mapOfReqs := make(map[string]map[string]struct{})
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		first := strings.TrimSpace(parts[0])
		second := strings.TrimSpace(parts[1])

		if _, exists := mapOfReqs[first]; !exists {
			mapOfReqs[first] = make(map[string]struct{})
		}
		mapOfReqs[first][second] = struct{}{}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		acc := make(map[string]struct{})
		items := strings.Split(line, ",")
		middleIndex := len(items) / 2
		correct := true

		for _, number := range items {
			reqs, _ := mapOfReqs[number]

			if existsInSet(acc, reqs) {
				correct = false
				break
			}

			acc[number] = struct{}{}
		}

		if correct {
			item, _ := strconv.Atoi(items[middleIndex])
			sum += item
		}
	}

	fmt.Println("Result: ", sum)
}
