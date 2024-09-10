package merge

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	it1 := NewListIterator([]string{"a", "d", "f"})
	it2 := NewListIterator([]string{"b", "e", "h"})
	it3 := NewListIterator([]string{"c", "g", "h"})

	it := NewMergeIterator([]Iterator{it1, it2, it3})
	var merged []string
	for it.Next() {
		merged = append(merged, it.At())
	}
	fmt.Println(merged)
}
