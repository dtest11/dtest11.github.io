package slidewindow

import "testing"
import "github.com/stretchr/testify/assert"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{"abcabcbb"}, want: 3},
		//{name: "2", args: args{"bbbbb"}, want: 1},
		//{name: "2", args: args{"pwwkew"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	t.Log(checkInclusion(s1, s2))

	s1 = "adc"
	s2 = "dcda"
	t.Log(checkInclusion(s1, s2))

	s1 = "abcdxabcde"
	s2 = "abcdeabcdx"
	assert.True(t, checkInclusion(s1, s2))
}
