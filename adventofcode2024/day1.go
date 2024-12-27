package adventofcode2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(inputFile string) ([]int, []int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	list1 := []int{}
	list2 := []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) < 2 {
			log.Fatalf("Invalid line format: %s", line)
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Error converting %s to integer: %v", parts[0], err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Error converting %s to integer: %v", parts[1], err)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list1, list2
}

func Day1(inputFile string) {
	// Get input from file
	list1, list2 := parseInput(inputFile)

	// Sort the lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Initial Value
	totalDistance := 0
	difference := 0
	similarityScoreTotal := 0

	// To share the same loop for performance
	for i := 0; i < len(list1); i++ {
		// part 1
		difference = list1[i] - list2[i]
		if difference < 0 {
			difference = 0 - difference
		}
		totalDistance += difference

		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				// part 2
				similarityScoreTotal += list1[i]
			}
		}
	}

	fmt.Printf("Part 1 %d\n", totalDistance)
	fmt.Printf("Part 2 %d\n", similarityScoreTotal)
}
