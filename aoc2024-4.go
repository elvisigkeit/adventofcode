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

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()

		if line == "" {
			break
		}

		parts := strings.Fields(line)

		safe := true
		firstFail := true
		decreasing := true
		first := true
		second := false
		previous := 0
		for _, part := range parts {
			num, _ := strconv.Atoi(part)

			if first {
				previous = num
				first = false
				second = true
				continue
			}
			if second {
				second = false
				decreasing = previous > num
			}
			if num == previous {
				if firstFail {
					firstFail = false
					previous = num
					continue
				} else {
					safe = false
					break
				}
			}

			delta := 0
			if decreasing {
				delta = previous - num
			} else {
				delta = num - previous
			}

			if delta <= 0 || delta > 3 {
				if firstFail {
					firstFail = false
				} else {
					safe = false
					break
				}
			}
			previous = num
		}

		if safe {
			sum += 1
		}
	}

	fmt.Println("The result is: ", sum)
}

// This code gets the wrong result, but by one, I had a typo and it was right
