package enumutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnum(t *testing.T) {
	as := assert.New(t)

	allowed := []string{"pending", "active", "closed"}
	defaultVal := "pending"

	e := NewEnum(allowed, defaultVal)

	// Test Initial Value
	as.Equal(defaultVal, e.String())
	as.Equal("string", e.Type())

	// Test Set Valid
	err := e.Set("active")
	as.NoError(err)
	as.Equal("active", e.String())

	// Test Set Invalid
	err = e.Set("unknown")
	as.Error(err)
	as.Equal("active", e.String()) // Value should not change
	as.Contains(err.Error(), "unknown is not included in")
}
