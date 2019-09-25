package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInputFromFile() [][]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var inputs [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var numbers []int
		for _, numberStr := range strings.Split(scanner.Text(), ", ") {
			number, _ := strconv.Atoi(numberStr)
			numbers = append(numbers, number)
		}
		inputs = append(inputs, numbers)
	}

	return inputs
}

func findPositiveMax(numbers []int) int {
	var positiveMax int
	for _, number := range numbers {
		if number > 0 && positiveMax < number {
			positiveMax = number
		}
	}
	return positiveMax
}

// Time complexity: 0(3n)
// Space complexity: Depends on max value in array
func lowestMissingPositiveInt(numbers []int) int {
	// 0(n)
	max := findPositiveMax(numbers)
	if max == 0 || max == 1 {
		return max + 1
	}
	keeper := make([]int, max+1)

	// O(n)
	for _, number := range numbers {
		if number > 0 {
			keeper[number] = number
		}
	}

	// 0(n)
	for i := 1; i < max+1; i++ {
		if keeper[i] == 0 {
			return i
		}
	}

	return max + 1
}

func main() {
	inputs := readInputFromFile()
	for _, numbers := range inputs {
		fmt.Printf("Numbers: %v\n", numbers)
		fmt.Printf("Lowest positive integer that does not exist in the array: %d\n\n", lowestMissingPositiveInt(numbers))
	}
}
