package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func existsInSet(set1, set2 map[string]struct{}) (bool, string) {
	for key := range set1 {
		if _, exists := set2[key]; exists {
			return true, key
		}
	}
	return false, ""
}

func makeCorrectItems(items []string, mapOfReqs map[string]map[string]struct{}) []string {
	acc := make(map[string]struct{})

	fmt.Println("Correcting items...")
	fmt.Println(items)
	for i := 0; i < len(items); i++ {
		number := items[i]
		reqs, _ := mapOfReqs[number]

		conflict, error := existsInSet(acc, reqs)
		acc[number] = struct{}{}
		if conflict {
			items = swap(items, error, number)
			acc = make(map[string]struct{})
			i = -1
		}
	}
	fmt.Println(items)

	return items
}

func swap(arr []string, first, second string) []string {
	index1, index2 := -1, -1
	for i, v := range arr {
		if v == first {
			index1 = i
		} else if v == second {
			index2 = i
		}
	}

	arr[index1], arr[index2] = arr[index2], arr[index1]
	return arr
}

// func reorder(arr []string, first, second string) []string {
// 	index1, index2 := -1, -1

// 	for i, v := range arr {
// 		if v == first {
// 			index1 = i
// 		} else if v == second {
// 			index2 = i
// 		}
// 	}

// 	// Remove the second value from its original position
// 	element := arr[index2]
// 	arr = append(arr[:index2], arr[index2+1:]...)

// 	newIndex := index1
// 	arr = append(arr[:newIndex], append([]string{element}, arr[newIndex:]...)...)

// 	return arr
// }

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

			conflict, _ := existsInSet(acc, reqs)
			if conflict {
				correct = false
				break
			}

			acc[number] = struct{}{}
		}

		if !correct {
			correctItems := makeCorrectItems(items, mapOfReqs)
			item, _ := strconv.Atoi(correctItems[middleIndex])
			sum += item
		}
	}

	fmt.Println("Result: ", sum)
}
