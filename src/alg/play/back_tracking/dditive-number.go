package back_tracking

import (
	"fmt"
	"strconv"
)

func isAdditiveNumber(num string) bool {
	var tmp []int64
	for _ , v := range []rune(num){
		fmt.Println(v)
		v1 , _ :=strconv.ParseInt(string(v),10,64)
		tmp =append(tmp,v1)
	}
	return isAdditiveNumberBackTrack(tmp, 2)
}

func isAdditiveNumberBackTrack(num []int64, curIndex int) bool {
	fmt.Println(curIndex-2,curIndex,"000000000000000000")
	if len(num) < 3 {
		return false
	}

	if curIndex >= len(num) {
		curIndex =len(num)-1
		third :=num[curIndex]
		for i := curIndex-1; i >= curIndex-2; i-- {
			fmt.Println(num[i],third,"111111")
			third-=num[i]

		}
		if third != 0 {
			return true
		}
		return  false
	}
	third := num[curIndex]
	for i := curIndex-1; i >= curIndex-2; i-- {
		fmt.Println(num[i],third)
		third -= num[i]
	}

	if third != 0 {
		return false
	}
	for i := curIndex; i < len(num); i++ {
		if isAdditiveNumberBackTrack(num, curIndex+2) {
			return false
		}
	}
	return true
}
