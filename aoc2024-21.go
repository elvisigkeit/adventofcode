package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func printLinkedList(ll *list.List) {
	for item := ll.Front(); item != nil; item = item.Next() {
		fmt.Printf("%d ", item.Value)
	}
	fmt.Println()
}

func main() {
	steps := 25
	ll := list.New()

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numbers := strings.Split(input, " ")
	for _, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		ll.PushBack(num)
	}

	for i := 0; i < steps; i++ {
		for item := ll.Front(); item != nil; item = item.Next() {
			value := item.Value
			if value == 0 {
				item.Value = 1
				continue
			}
			var digits []int
			currentValue := item.Value.(int)
			for currentValue > 0 {
				digit := currentValue % 10
				digits = append(digits, digit)
				currentValue /= 10
			}
			slices.Reverse(digits)
			if len(digits)%2 == 0 {
				leftHalf := digits[:len(digits)/2]
				rightHalf := digits[len(digits)/2:]
				newLeft := 0
				for _, digit := range leftHalf {
					newLeft = newLeft*10 + digit
				}
				newRight := 0
				for _, digit := range rightHalf {
					newRight = newRight*10 + digit
				}
				ll.InsertBefore(newLeft, item)
				item.Value = newRight
				continue
			}
			item.Value = item.Value.(int) * 2024
		}

		// printLinkedList(ll)
		fmt.Println("Step: ", i)
		fmt.Println("Count: ", ll.Len())
	}

	fmt.Println("\nCount: ", ll.Len())
}
