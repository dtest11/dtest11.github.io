package bit

import (
	"fmt"
	"testing"
)

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	find := func(w string) bool {
		for _, ws := range w {
			found :=false
			for _, v := range allowed {
				if ws == v {
					found=true
					break
				}
			}
			if !found{
				return false
			}
		}
		return  true
	}
	for _, w := range words {
		if find(w){
			count++
		}
	}
	return count
}

func Test_count(t *testing.T) {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	a :=countConsistentStrings(allowed, words)
	fmt.Println(a)
}
