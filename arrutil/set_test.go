package arrutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	as := assert.New(t)

	// Test NewSet
	s := NewSet("a", "b", "c")
	as.Equal(3, len(s))
	as.True(s.Has("a"))
	as.True(s.Has("b"))
	as.True(s.Has("c"))
	as.False(s.Has("d"))

	// Test Add
	s.Add("d")
	as.True(s.Has("d"))
	as.Equal(4, len(s))

	// Test Delete
	s.Delete("a")
	as.False(s.Has("a"))
	as.Equal(3, len(s))

	// Test Duplicate Add (should not change length)
	s.Add("b")
	as.Equal(3, len(s))
}
