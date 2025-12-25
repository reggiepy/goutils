package arrutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInArray(t *testing.T) {
	as := assert.New(t)

	// Test Slice
	slice := []int{1, 2, 3, 4, 5}
	as.True(InArray(3, slice))
	as.False(InArray(6, slice))

	// Test Array
	array := [3]string{"a", "b", "c"}
	as.True(InArray("b", array))
	as.False(InArray("d", array))

	// Test Map (checking values)
	m := map[string]int{"a": 1, "b": 2}
	as.True(InArray(1, m))
	as.False(InArray(3, m))

	// Test Invalid Type (Panic)
	as.Panics(func() {
		InArray(1, "not an array")
	})
}
