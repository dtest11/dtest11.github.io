package pow

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myPow(t *testing.T) {
	res := fastMul(5, 10)
	assert.Equal(t, math.Pow(5, 10), res)

	res = fastMul(2.1, 3)
	assert.Equal(t, math.Pow(2.1, 3), res)

	res = fastMul(2.00, -2) * 1.0
	assert.Equal(t, math.Pow(2.00, -2), res)
}
