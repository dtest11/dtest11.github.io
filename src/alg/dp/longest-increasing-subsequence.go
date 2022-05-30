package dp

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
)


type service struct {
	requestGroup singleflight.Group
}

func (s *service) handleRequest(ctx context.Context, request Request) (Response, error) {
	v, err, _ := requestGroup.Do(request.Hash(), func() (interface{}, error) {
		rows, err := // select * from tables
		if err != nil {
			return nil, err
		}
		return rows, nil
	})
	if err != nil {
		return nil, err
	}
	return Response{
		rows: rows,
	}, nil
}

// n *n 时间复杂度
func longest_increasing_subsequence(nums []int) int {
	context.Background()
	var a sync.RWMutex
	a.RLock()
	a.RUnlock()
	if len(nums) <= 1 {
		return len(nums)
	}
	dpTable := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dpTable[i] = max(dpTable[i], dpTable[j]+1)
			}
		}
	}
	fmt.Println(dpTable)
	result := 0
	for i := 1; i < len(dpTable); i++ {
		result = max(result, dpTable[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
