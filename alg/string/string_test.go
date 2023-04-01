package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	str := "III"
	assert.Equal(t, 3, romanToInt(str))

	str = "LVIII"
	assert.Equal(t, 58, romanToInt(str))

	str = "MCMXCIV"
	assert.Equal(t, 1994, romanToInt(str))
	res := isPalindrome("A man, a plan, a canal: Panama")
	//res := isPalindrome("aa")
	t.Log(res)

}
