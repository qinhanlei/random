package random

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

const (
	TOTAL_TEST_TIMES = 10
)

const (
	NUM_LEN            = 10
	NUM_RANGE          = 100
	SHUFFLE_TIMES      = 10000 * NUM_LEN
	TOLERANT_PRECISION = 0.01 //NOTE: just ok for above, have no idea to found formula
)

func TestSortStringSlice(t *testing.T) {
	errCnt := 0
	fmt.Printf("tolerant precision: %v\n", TOLERANT_PRECISION)

	for testId := 0; testId < TOTAL_TEST_TIMES; testId++ {
		var nums IntSlice = make([]int, NUM_LEN)
		for i, _ := range nums {
			nums[i] = rand.Int() % NUM_RANGE
		}
		fmt.Println("init nums:", nums)

		sum := 0
		for _, v := range nums {
			sum += v
		}
		expectVal := sum * (SHUFFLE_TIMES / NUM_LEN)
		fmt.Println("every bucket expect value:", expectVal)

		buckets := make(map[int]int)
		for counter := 0; counter < SHUFFLE_TIMES; counter++ {
			nums.Shuffle()
			for i, v := range nums {
				buckets[i] += v
			}
		}

		for k, v := range buckets {
			d := v - expectVal
			p := float64(d) / float64(expectVal)
			fmt.Printf("bucket:%v value:%v deviation:%v, %v\n", k, v, d, p)
			if math.Abs(p) > TOLERANT_PRECISION {
				errCnt++
			}
		}
	}

	fmt.Println("errCnt:", errCnt)
	if errCnt != 0 {
		t.Errorf("over tolerant precision %v times.", errCnt)
	}
}
