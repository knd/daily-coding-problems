package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInputFromFile() ([]int, int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var numbers []int
	var targetSum int

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, numberStr := range strings.Split(scanner.Text(), " ") {
		number, _ := strconv.Atoi(numberStr)
		numbers = append(numbers, number)
	}
	scanner.Scan()
	targetSum, _ = strconv.Atoi(scanner.Text())

	return numbers, targetSum
}

func result(numbers []int, targetSum int) bool {
	difference := map[int]bool{}

	for _, number := range numbers {
		if _, ok := difference[number]; ok {
			return true
		}
		difference[targetSum-number] = true
	}

	return false
}

func main() {
	numbers, targetSum := readInputFromFile()
	fmt.Printf("Any 2 numbers from list add up to k: %t\n", result(numbers, targetSum))
}
