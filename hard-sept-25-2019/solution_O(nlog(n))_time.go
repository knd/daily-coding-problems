package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

// Time complexity: O(n*log(n))
// Space complexity: O(1)
func lowestMissingPositiveInt(numbers []int) int {
	sort.Ints(numbers) // O(n*log(n))

	// O(n)
	var positivePrev int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 && positivePrev == 0 {
			positivePrev = numbers[i]
		} else if positivePrev > 0 {
			if positivePrev+1 < numbers[i] {
				return positivePrev + 1
			}
			positivePrev = numbers[i]
		}
	}

	return positivePrev + 1
}

func main() {
	inputs := readInputFromFile()
	for _, numbers := range inputs {
		fmt.Printf("Numbers: %v\n", numbers)
		fmt.Printf("Lowest positive integer that does not exist in the array: %d\n\n", lowestMissingPositiveInt(numbers))
	}
}
