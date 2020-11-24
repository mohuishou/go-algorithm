package stack

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	assert.Equal(t, "", stack.Pop())
	stack.Push("1")
	stack.Push("2")
	assert.Equal(t, "2", stack.Pop())
	assert.Equal(t, "1", stack.Pop())
	for i := 3; i < 20; i++ {
		stack.Push(strconv.Itoa(i))
	}
	assert.Equal(t, "19", stack.Pop())
}
