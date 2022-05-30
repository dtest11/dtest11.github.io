package bubble

import (
	"golang.org/x/exp/constraints"

	"fmt"
)

func Bubble[T constraints.Ordered](data []T) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			compare := data[j] < data[j-1]
			if compare {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
		fmt.Println("sort round:", data)
	}
}
