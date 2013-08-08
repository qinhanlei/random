package random

import (
	//"fmt"
	"math/rand"
	"time"
)

func init() {
	// just for convenient
	rand.Seed(time.Now().UnixNano())
}

//NOTE: from golang's official sort pkg below
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func Shuffle(data Interface) {
	n := data.Len()
	//NOTE: from gcc4.7.1's STL source code
	for i := 1; i < n; i++ {
		data.Swap(i, rand.Int()%(i+1))
	}
}

//NOTE: from golang's official sort pkg below
type IntSlice []int

func (p IntSlice) Len() int      { return len(p) }
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p IntSlice) Shuffle()      { Shuffle(p) }
