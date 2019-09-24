package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Node represents a node in a tree
type Node struct {
	left  *Node
	right *Node
	value string
}

const NilValue string = "NIL" // change this to special value

// Serialize returns serialized string of tree
func serialize(root *Node) string {
	if root == nil {
		return NilValue
	}

	return fmt.Sprintf("%s %s %s", root.value, serialize(root.left), serialize(root.right))
}

// Deserialize returns binary tree from serialized form
func deserialize(serializedTree string) *Node {
	tokens := tokenize(serializedTree)

	root, _ := deserializeFromTokens(tokens, 0)
	return root
}

func deserializeFromTokens(tokens []string, index int) (*Node, int) {
	if index >= len(tokens) || tokens[index] == NilValue {
		return nil, index
	}

	root := &Node{value: tokens[index]}
	root.left, index = deserializeFromTokens(tokens, index+1)
	root.right, index = deserializeFromTokens(tokens, index+1)

	return root, index
}

func tokenize(serializedTree string) []string {
	return strings.Split(serializedTree, " ")
}

func TestFullTree(t *testing.T) {
	assert := assert.New(t)
	root := &Node{
		value: "root",
		left: &Node{
			value: "left",
			left: &Node{
				value: "left.left",
			},
			right: &Node{
				value: "left.right",
			},
		},
		right: &Node{
			value: "right",
			left: &Node{
				value: "right.left",
			},
			right: &Node{
				value: "right.right",
			},
		},
	}

	serializedTree := serialize(root)
	fmt.Printf("Serialized tree: %s\n", serializedTree)

	t.Run("root has correct value", func(t *testing.T) {
		assert.Equal("root", deserialize(serializedTree).value)
	})

	t.Run("root's left has correct value", func(t *testing.T) {
		assert.Equal("left", deserialize(serializedTree).left.value)
	})

	t.Run("root's right has correct value", func(t *testing.T) {
		assert.Equal("right", deserialize(serializedTree).right.value)
	})

	t.Run("left's left has correct value", func(t *testing.T) {
		assert.Equal("left.left", deserialize(serializedTree).left.left.value)
	})

	t.Run("left's right has correct value", func(t *testing.T) {
		assert.Equal("left.right", deserialize(serializedTree).left.right.value)
	})

	t.Run("right's left has correct value", func(t *testing.T) {
		assert.Equal("right.left", deserialize(serializedTree).right.left.value)
	})

	t.Run("right's right has correct value", func(t *testing.T) {
		assert.Equal("right.right", deserialize(serializedTree).right.right.value)
	})

	t.Run("left's left is leave node", func(t *testing.T) {
		assert.Nil(deserialize(serializedTree).left.left.left)
		assert.Nil(deserialize(serializedTree).left.left.right)
	})

	t.Run("left's right is leave node", func(t *testing.T) {
		assert.Nil(deserialize(serializedTree).left.right.left)
		assert.Nil(deserialize(serializedTree).left.right.right)
	})

	t.Run("right's left is leave node", func(t *testing.T) {
		assert.Nil(deserialize(serializedTree).right.left.left)
		assert.Nil(deserialize(serializedTree).right.left.right)
	})

	t.Run("right's right is leave node", func(t *testing.T) {
		assert.Nil(deserialize(serializedTree).right.right.left)
		assert.Nil(deserialize(serializedTree).right.right.right)
	})
}

func TestUnbalancedTree(t *testing.T) {
	assert := assert.New(t)
	root := &Node{
		value: "root",
		left: &Node{
			value: "left",
		},
		right: &Node{
			value: "right",
			right: &Node{
				value: "right.right",
			},
		},
	}

	serializedTree := serialize(root)
	fmt.Printf("Serialized tree: %s\n", serializedTree)

	t.Run("root has correct value", func(t *testing.T) {
		assert.Equal("root", deserialize(serializedTree).value)
	})

	t.Run("root's left has correct value", func(t *testing.T) {
		assert.Equal("left", deserialize(serializedTree).left.value)
	})

	t.Run("root's right has correct value", func(t *testing.T) {
		assert.Equal("right", deserialize(serializedTree).right.value)
	})

	t.Run("left is leave node", func(t *testing.T) {
		assert.Nil(deserialize(serializedTree).left.left)
		assert.Nil(deserialize(serializedTree).left.right)
	})

	t.Run("right has no left node", func(t *testing.T) {
		assert.Nil(deserialize(serializedTree).right.left)
	})

	t.Run("right's right has correct value", func(t *testing.T) {
		assert.Equal("right.right", deserialize(serializedTree).right.right.value)
	})

	t.Run("right's right is leave node", func(t *testing.T) {
		assert.Nil(deserialize(serializedTree).right.right.left)
		assert.Nil(deserialize(serializedTree).right.right.right)
	})
}

func TestSingleNodeTree(t *testing.T) {
	assert := assert.New(t)
	root := &Node{value: "root"}

	serializedTree := serialize(root)
	fmt.Printf("Serialized tree: %s\n", serializedTree)

	t.Run("root has correct value", func(t *testing.T) {
		assert.Equal("root", deserialize(serializedTree).value)
	})

	t.Run("root is leave node", func(t *testing.T) {
		assert.Nil(deserialize(serializedTree).left)
		assert.Nil(deserialize(serializedTree).right)
	})
}
