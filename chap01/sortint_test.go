package chap01

import (
	"github.com/cyfdecyf/goutil/rand"
	"sort"
	"testing"
)

func testSortInt(n, max int, t *testing.T) {
	if n > max {
		t.Fatal("testSortInt: n > max")
	}
	arr := rand.GenKRandomLessN(n, max)
	SortInt(arr, max)

	if !sort.IsSorted(sort.IntSlice(arr)) {
		t.Errorf("SortInt error %d intergers less than %d\n", n, max)
	}
}

func TestSortInt(t *testing.T) {
	testSortInt(100, 1000, t)
	testSortInt(1000, 1001, t)
}
