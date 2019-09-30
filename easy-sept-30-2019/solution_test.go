package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Node represents a node in a tree
type Node struct {
	left  *Node
	right *Node
	value int
}

func isUnival(root *Node) bool {
	if root.left == nil && root.right == nil {
		return true
	} else if root.left == nil {
		return isUnival(root.right)
	} else if root.right == nil {
		return isUnival(root.left)
	}

	return root.left.value == root.right.value &&
		isUnival(root.left) &&
		isUnival(root.right)
}

// Only benefits if tree has high depth
func isUnivalWithMemoization(root *Node, univalMemoization *map[*Node]bool) bool {
	if val, ok := (*univalMemoization)[root]; ok && val {
		return true
	}

	if root.left == nil && root.right == nil {
		(*univalMemoization)[root] = true
		return true
	} else if root.left == nil {
		return isUnivalWithMemoization(root.right, univalMemoization)
	} else if root.right == nil {
		return isUnivalWithMemoization(root.left, univalMemoization)
	}

	unival := root.left.value == root.right.value &&
		isUnivalWithMemoization(root.left, univalMemoization) &&
		isUnivalWithMemoization(root.right, univalMemoization)

	(*univalMemoization)[root] = true

	return unival
}

func preOrderNodes(root *Node, stack *[]*Node) {
	if root == nil {
		return
	}
	preOrderNodes(root.left, stack)
	*stack = append(*stack, root)
	preOrderNodes(root.right, stack)
}

func printTimeStats(start time.Time, end time.Time) {
	elapsed := (end.Sub(start)).Nanoseconds()
	fmt.Printf("Elapsed time: %d ns\n", elapsed)
}

func univalCount(root *Node) int {
	startTime := time.Now()
	var count int
	var nodeStack []*Node
	univalMemoization := map[*Node]bool{}
	preOrderNodes(root, &nodeStack)

	for _, node := range nodeStack {
		// Replace isUnivalWithMemoization with isUnival for performance testing
		if isUnivalWithMemoization(node, &univalMemoization) {
			count++
		}
	}

	endTime := time.Now()
	printTimeStats(startTime, endTime)
	return count
}

func Test_FullTree(t *testing.T) {
	assert := assert.New(t)
	root := &Node{
		value: 0,
		left:  &Node{value: 1},
		right: &Node{
			value: 5,
			left: &Node{
				value: 1,
				left:  &Node{value: 1},
				right: &Node{value: 1},
			},
			right: &Node{value: 0},
		},
	}

	assert.Equal(5, univalCount(root))
}

func Test_SingleNodeTree(t *testing.T) {
	assert := assert.New(t)
	root := &Node{value: 1}

	assert.Equal(1, univalCount(root))
}

func Test_TriadNodeTree(t *testing.T) {
	assert := assert.New(t)
	root := &Node{
		value: 0,
		left:  &Node{value: 1},
		right: &Node{value: 1},
	}

	assert.Equal(3, univalCount(root))
}

func Test_TriadNodeTreeAllSameValues(t *testing.T) {
	assert := assert.New(t)
	root := &Node{
		value: 1,
		left:  &Node{value: 1},
		right: &Node{value: 1},
	}

	assert.Equal(3, univalCount(root))
}

func Test_DeepBinaryTree(t *testing.T) {
	assert := assert.New(t)
	root := &Node{
		value: 1,
		right: &Node{
			value: 1,
			left: &Node{
				value: 1,
				left:  &Node{value: 1},
				right: &Node{value: 1},
			},
			right: &Node{value: 1},
		},
	}

	assert.Equal(6, univalCount(root))
}

func Test_BiNodeTree(t *testing.T) {
	assert := assert.New(t)
	root := &Node{
		value: 1,
		right: &Node{value: 1},
	}

	assert.Equal(2, univalCount(root))
}

func Test_BinNodeDifferentValues(t *testing.T) {
	assert := assert.New(t)
	root := &Node{
		value: 1,
		right: &Node{value: 0},
	}

	assert.Equal(2, univalCount(root))
}
