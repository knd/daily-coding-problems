package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInputFromFile() []int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, numberStr := range strings.Split(scanner.Text(), " ") {
		number, _ := strconv.Atoi(numberStr)
		numbers = append(numbers, number)
	}
	scanner.Scan()

	return numbers
}

func result(numbers []int) []int {
	if len(numbers) <= 1 {
		panic("Length of given array must be 2 or more")
	}

	// [1, 2, 3, 4, 5] -> [1, 2, 6, 24, 120]
	forwardProducts := make([]int, len(numbers))
	forwardProducts[0] = numbers[0]
	for i := 1; i < len(numbers); i++ {
		forwardProducts[i] = forwardProducts[i-1] * numbers[i]
	}

	// [1, 2, 3, 4, 5] -> [120, 120, 60, 20, 5]
	backwardProducts := make([]int, len(numbers))
	backwardProducts[len(numbers)-1] = numbers[len(numbers)-1]
	for i := len(numbers) - 2; i >= 0; i-- {
		backwardProducts[i] = backwardProducts[i+1] * numbers[i]
	}

	products := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			products[i] = backwardProducts[i+1]
		} else if i == len(numbers)-1 {
			products[i] = forwardProducts[i-1]
		} else {
			products[i] = forwardProducts[i-1] * backwardProducts[i+1]
		}
	}

	return products
}

func main() {
	numbers := readInputFromFile()
	fmt.Printf("Given array:\n%v\n", numbers)
	fmt.Printf("New array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i:\n%v\n", result(numbers))
}
