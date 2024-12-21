package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	var disk []int

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()

		if line == "" {
			fmt.Println("Calculating...")
			break
		}

		for _, ch := range line {
			disk = append(disk, int(ch)-48)
		}
	}

	j := len(disk) - 1
	if j%2 == 1 {
		j--
	}

	di := 0
	for i := 0; i <= j; i++ {
		if i%2 == 0 {
			for disk[i] > 0 {
				id := i / 2
				sum += di * id
				di++
				disk[i]--
				// fmt.Print(id)
				// fmt.Printf(": %d * %d = %d\n", di, id, sum)
			}
		} else {
			for disk[i] > 0 {
				if disk[j] <= 0 {
					j -= 2
				}
				id := j / 2
				sum += di * id
				di++
				disk[i]--
				disk[j]--
				if disk[j] <= 0 {
					j -= 2
				}
				// fmt.Print(id)
				// fmt.Printf(": %d * %d = %d\n", di, id, sum)
			}
		}
		// fmt.Println(sum)
	}

	fmt.Println("\nThe result is: ", sum)
}
