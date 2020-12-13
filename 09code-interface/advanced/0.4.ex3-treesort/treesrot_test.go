package treesort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestTree(t *testing.T) {
	data := make([]int, 50)

	for i := range data {
		data[i] = rand.Int() % 50
	}

	fmt.Printf("before %#v\n", data)

	Sort(data)

	fmt.Printf("after %#v\n", data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("data not sort")
	}
}

func TestString(t *testing.T) {
	tree := &tree{0, nil, nil}

	add(tree, 1)
	add(tree, 2)

	fmt.Printf("%#v\n", tree)
	
	want := "012"
	got := tree.String()

	if got != want {
		t.Fatalf("want: %s, got %s\n", want, got)
	}
}
