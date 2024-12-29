package adventofcode2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getOutputFromMultIns(ins string, re *regexp.Regexp) int {
	matches := re.FindAllStringSubmatch(ins, -1)
	if matches != nil {
		num1, err := strconv.Atoi(matches[0][1])
		if err != nil {
			log.Printf("Error converting %d to integer: %v", num1, err)
			num1 = 0
		}

		num2, err := strconv.Atoi(matches[0][2])
		if err != nil {
			log.Printf("Error converting %d to integer: %v", num2, err)
			num2 = 0
		}
		return num1 * num2
	}
	return 0
}

func getInstructionsFromFile(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`(?m)mul\(\d+,\d+\)|don\'t\(\)|do\(\)`)

	var instructions []string

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		if matches != nil {
			instructions = append(instructions, matches...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return instructions
}

func Day3(inputFile string) {
	instructions := getInstructionsFromFile(inputFile)

	on := true
	var mult, mults, multsConditional int
	re := regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)`)

	for _, instruction := range instructions {
		switch instruction {
		case "do()":
			on = true
		case "don't()":
			on = false
		default:
			mult = getOutputFromMultIns(instruction, re)
			mults += mult
			if on {
				multsConditional += mult
			}
		}
	}
	fmt.Printf("Part 1: %v\n", mults)
	fmt.Printf("Part 2: %v\n", multsConditional)
}
