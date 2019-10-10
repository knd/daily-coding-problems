package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func readInputFromFile() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var dict []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Split(line, ",") {
			dict = append(dict, strings.TrimSpace(word))
		}
	}

	return dict
}

// Node represents a character
type Node struct {
	char     rune
	children []*Node
}

func buildWordTrie(root *Node, word string, index int) {
	if index >= len(word) {
		return
	}

	for _, child := range root.children {
		if child.char == rune(word[index]) {
			buildWordTrie(child, word, index+1)
			return
		}
	}

	child := &Node{char: rune(word[index])}
	root.children = append(root.children, child)
	buildWordTrie(child, word, index+1)
}

func buildTrie(dict []string) *Node {
	trie := &Node{}
	for _, word := range dict {
		buildWordTrie(trie, word, 0)
	}
	return trie
}

func printTrie(node *Node, depth int) {
	for i := 1; i < depth; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("%s\n", string(node.char))
	for _, child := range node.children {
		printTrie(child, depth+1)
	}
}

func getPivotNode(word string, index int, children []*Node) *Node {
	if index >= len(word) {
		return nil
	}
	currChar := rune(word[index])
	for _, child := range children {
		if currChar == (*child).char {
			if index == len(word)-1 {
				return child
			}
			return getPivotNode(word, index+1, child.children)
		}
	}
	return nil
}

func getRemainingWords(children []*Node, temp string, result *[]string) {
	if len(children) == 0 {
		*result = append(*result, temp)
		return
	}
	for _, child := range children {
		getRemainingWords(child.children, temp+string((*child).char), result)
	}
}

func recommendedWords(word string, trie *Node) []string {
	startTime := time.Now()

	var result []string
	pivotNode := getPivotNode(word, 0, trie.children)
	var remainingWords []string
	if pivotNode != nil {
		getRemainingWords(pivotNode.children, "", &remainingWords)
	}

	for _, rm := range remainingWords {
		result = append(result, word+rm)
	}

	endTime := time.Now()
	elapsed := endTime.Sub(startTime).Nanoseconds()
	fmt.Printf("Elapsed time: %dns\n", elapsed)

	return result
}

func naive(word string, dict []string) []string {
	startTime := time.Now()

	var result []string
	for _, w := range dict {
		if strings.HasPrefix(w, word) {
			result = append(result, w)
		}
	}

	endTime := time.Now()
	elapsed := endTime.Sub(startTime).Nanoseconds()
	fmt.Printf("Elapsed time: %dns\n", elapsed)

	return result
}

func main() {
	dict := readInputFromFile()

	// trie := buildTrie(dict)
	// for {
	// 	var word string
	// 	fmt.Scanf("%s", &word)
	// 	fmt.Printf("You mean: %v\n", recommendedWords(word, trie))
	// }

	for {
		var word string
		fmt.Scanf("%s", &word)
		fmt.Printf("You mean: %v\n", naive(word, dict))
	}
}
