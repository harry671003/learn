package merge

import (
	"container/heap"
	"strings"
)

/*
// An iterator implements the following interface.
interface Iterator {
	string At()
	bool Next()
}

// Assume we have a list iterator implementation that iterates over a list.
List l1 = Arrays.asList("a", "c", "g");
List l2 = Arrays.asList("b", "f", "h");
List l3 = Arrays.asList("d", "e", "f");

Iterator it1 = new ListIterator(l1)
Iterator it2 = new ListIterator(l2)
Iterator it3 = new ListIterator(l3)

// The values can be read as follows.
it1.Next() 			// true
it1.At() 			// "a"
it1.Next() 			// true
it1.At() 			// "b"
it1.Next() 			// true
it1.At() 			// "c"
it1.Next() 			// false
it1.At()	 		// error


// Given a list of N such sorted iterators implement another iterator that merges them.
List its = Arrays.asList(it1, it2, it3)
Iterator m = new MergeIterator(its)

while m.Next() {
	System.out.println(m.At())
}
// output: abcdeffgh
*/

type Iterator interface {
	At() string
	Next() bool
}

// An ItHeap is a min-heap of ints.
type IteratorHeap []Iterator

func (h IteratorHeap) Len() int {
	return len(h)
}
func (h IteratorHeap) Less(i, j int) bool {
	a := h[i].At()
	b := h[j].At()
	return strings.Compare(a, b) < 0
}
func (h IteratorHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IteratorHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Iterator))
}

func (h *IteratorHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewListIterator(list []string) *ListIterator {
	return &ListIterator{
		list: list,
	}
}

// ListIterator converts a list into iterator.
// Use this for demoing that
type ListIterator struct {
	list []string
	cur  int
}

func (it *ListIterator) At() string {
	return it.list[it.cur-1]
}

func (it *ListIterator) Next() bool {
	it.cur++
	return it.cur-1 < len(it.list)
}

// MergeIterator merges multiple iterators.
// Duplicates are returned.
type MergeIterator struct {
	h   IteratorHeap
	cur Iterator
}

func NewMergeIterator(its []Iterator) Iterator {
	h := IteratorHeap{}

	for _, it := range its {
		if it.Next() {
			heap.Push(&h, it)
		}
	}
	return &MergeIterator{
		h: h,
	}
}

func (m *MergeIterator) At() string {
	return m.cur.At()
}

func (m *MergeIterator) Next() bool {
	// Increment current
	if m.cur != nil {
		if m.cur.Next() {
			heap.Push(&m.h, m.cur)
		}
	}

	if len(m.h) == 0 {
		return false
	}

	// This will be the new current
	m.cur = heap.Pop(&m.h).(Iterator)
	return true
}
