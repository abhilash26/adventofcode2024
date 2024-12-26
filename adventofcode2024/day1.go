package adventofcode2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sort(list *[]int) {
	length := len(*list)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if (*list)[j] > (*list)[j+1] {
				(*list)[j], (*list)[j+1] = (*list)[j+1], (*list)[j]
			}
		}
	}
}

func Day1(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	list1 := []int{}
	list2 := []int{}
	totalDistance := 0

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

	sort(&list1)
	sort(&list2)

	// Insertion sort can improve performance, when used within the loop

	difference := 0
	similarityScoreTotal := 0

	for i := 0; i < len(list1); i++ {

		difference = list1[i] - list2[i]
		if difference < 0 {
			difference = 0 - difference
		}
		totalDistance += difference

		// To share the same loop for performance
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				similarityScoreTotal += list1[i]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 %d\n", totalDistance)
	fmt.Printf("Part 2 %d\n", similarityScoreTotal)
}
