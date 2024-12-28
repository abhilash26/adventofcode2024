package adventofcode2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func getInt(a string) int {
	num, err := strconv.Atoi(a)
	if err != nil {
		log.Printf("Error converting %s to integer: %v", a, err)
		return 0
	}
	return num
}

func checkAscending(l []int) bool {
	for i := 0; i < len(l)-1; i++ {
		d := l[i+1] - l[i]
		if d < 1 || d > 3 {
			return false
		}
	}
	return true
}

func checkDecending(l []int) bool {
	for i := 0; i < len(l)-1; i++ {
		d := l[i] - l[i+1]
		if d < 1 || d > 3 {
			return false
		}
	}
	return true
}

func isDampened(l []int) bool {
	length := len(l)
	for i := 0; i < length; i++ {
		newSlice := removeIndex(l, i)
		if isSafe(newSlice) {
			return true
		}
	}
	return false
}

func isSafe(l []int) bool {
	return checkAscending(l) || checkDecending(l)
}

func Day2(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safeSum := 0
	safeSumDampener := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := []int{}
		line := scanner.Text()
		parts := strings.Fields(line)

		for _, part := range parts {
			levels = append(levels, getInt(part))
		}

		if isSafe(levels) {
			safeSum++
			safeSumDampener++
		} else {
			if isDampened(levels) {
				safeSumDampener++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", safeSum)
	fmt.Printf("Part 2: %d\n", safeSumDampener)
}
