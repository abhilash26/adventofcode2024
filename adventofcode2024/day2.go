package adventofcode2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isDampened(l []int) bool {
	length := len(l)
	for i := 0; i < length; i++ {
		// Create a new slice with the item removed by index
		newSlice := make([]int, 0, length-1)
		newSlice = append(newSlice, l[:i]...)
		newSlice = append(newSlice, l[i+1:]...)

		if isSafe(newSlice) {
			return true
		}
	}
	return false
}

func isSafe(l []int) bool {
	length := len(l)

	if length < 2 {
		return false
	}

	d := l[1] - l[0]
	if d == 0 || (d < -3 || d > 3) {
		return false
	}

	isAscending := d > 0

	for i := 1; i < length-1; i++ {
		d = l[i+1] - l[i]
		if isAscending {
			if d < 1 || d > 3 {
				return false
			}
		} else {
			if d > -1 || d < -3 {
				return false
			}
		}
	}

	return true
}

func Day2(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var safeSum, safeSumDampener int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Printf("Error converting %s to integer: %v", part, err)
				num = 0
			}
			levels[i] = num
		}

		if isSafe(levels) {
			safeSum++
			safeSumDampener++
		} else if isDampened(levels) {
			safeSumDampener++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", safeSum)
	fmt.Printf("Part 2: %d\n", safeSumDampener)
}
