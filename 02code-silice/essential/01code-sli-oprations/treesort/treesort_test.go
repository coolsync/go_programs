package treesort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 10)
	for i := range data {
		data[i] = rand.Int() % 10
	}
	fmt.Printf("before: %#v\n", data)
	Sort(data)
	fmt.Printf("after: %#v\n", data)

	if !sort.IntsAreSorted(data) {
		t.Error("data not sorted!")
	}
}