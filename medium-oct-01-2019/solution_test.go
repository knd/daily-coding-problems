package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func actualCount(message string, pivotMsg string) int {
	if len(message) == 0 {
		if len(pivotMsg) > 0 {
			return 1
		}
		return 0
	}
	if len(message) == 1 {
		return 1
	}

	count := 0
	for i := 0; i < len(message); i++ {
		pivotMsg := message[:(i + 1)]
		num, _ := strconv.Atoi(pivotMsg)
		if num > 26 {
			break
		}
		count += actualCount(message[(i+1):len(message)], pivotMsg)
	}

	return count
}

func decodeStyleCount(message string) int {
	return actualCount(message, "")
}

func Test_EmptyStr(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, decodeStyleCount(""))
}

func Test_EdgeCase1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, decodeStyleCount("1"))
}

func Test_111(t *testing.T) {
	assert := assert.New(t)
	// 1, 1, 1 | 11, 1 | 1, 11
	assert.Equal(3, decodeStyleCount("111"))
}

func Test_123(t *testing.T) {
	assert := assert.New(t)
	// 1, 2, 3 | 12, 3 | 1, 23 |
	assert.Equal(3, decodeStyleCount("123"))
}

func Test_26(t *testing.T) {
	assert := assert.New(t)
	// 2, 6 | 26
	assert.Equal(2, decodeStyleCount("26"))
}

func Test_12345(t *testing.T) {
	assert := assert.New(t)
	// 1, 2, 3, 4, 5 || 1, 23, 4, 5
	// 12, 3, 4, 5
	assert.Equal(3, decodeStyleCount("12345"))
}
