package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	source := [3]string{"hello", "world", "abac"}
	output := [3]string{"holle", "world", "abac"}
	for index, val := range source {
		result := ReverseVowels(val)
		assert.Equal(t, result, output[index])
	}
}
