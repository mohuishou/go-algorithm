package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLRU(t *testing.T) {
	l := NewLRU(3)
	assert.Equal(t, l.Get(""), nil)

	l.Put("1", 1)
	l.Put("2", 2)
	l.Put("3", 3)
	assert.Equal(t, 3, l.Get("3"))
	assert.Equal(t, 1, l.Get("1"))

	l.Put("4", 4)
	assert.Equal(t, nil, l.Get("2"))

	l.Put("3", 31)
	assert.Equal(t, 31, l.Get("3"))
}
