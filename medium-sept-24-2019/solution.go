package main

import (
	"fmt"
	"testing"

	"github.com/knd/daily-coding-problems/medium-sept-24-2019/node"
	"github.com/stretchr/testify/assert"
)

func main() {
	leftNode := &node.Node{left: nil, right: nil, value: "leftNode"}
	rightNode := &node.Node{left: nil, right: nil, value: "rightNode"}
	rootNode := &node.Node{left: leftNode, right: rightNode, value: "root"}

	fmt.Printf("Root=%s\n", rootNode.value)
	fmt.Printf("Left=%s\n", rootNode.left.value)
	fmt.Printf("Right=%s\n", rootNode.right.value)
}

func testFullTree(t *testing.T) {
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

func testUnbalancedTree(t *testing.T) {
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

func testSingleNodeTree(t *testing.T) {
	assert := assert.New(t)
	root := &Node{value: "root"}

	serializedTree := serialize(root)

	t.Run("root has correct value", func(t *testing.T) {
		assert.Equal("root", deserialize(serializedTree).value)
	})

	t.Run("root is leave node", func(t *testing.T) {
		assert.Nil(deserialize(serializedTree).left)
		assert.Nil(deserialize(serializedTree).right)
	})
}
