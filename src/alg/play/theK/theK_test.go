package thek

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}
func TestFindKthLargest(t *testing.T) {
	data :=pre()
	FindKthLargest(data,6)
	fmt.Println("-------")
	QuickSort(data)
	fmt.Println(data)

}

func del(data []int)  {
	data[0]=10
}

func TestQuickSort(t *testing.T) {
    data :=pre()
    QuickSort(data)
    //del(data)
    fmt.Println(data)

}

func pre() []int {
	var data []int
	for i := 0; i < 10; i++ {
		data = append(data, rand.Intn(100))
	}
	log.Println("data: ",data)

	return data
}
