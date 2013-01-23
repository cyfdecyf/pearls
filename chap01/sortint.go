// This package implements the integer sorting algorithm in chapter 1.
package chap01

import (
	"github.com/cyfdecyf/goutil/bitset"
)

// SortInt sorts slice arr which contains integers that are unique and
// less than max.
// It panics if len(a) >= max.
func SortInt(arr []int, max int) {
	n := len(arr)
	if n >= max {
		panic("SortInt input array invalid: length > max integer")
	}

	bs := bitset.NewBitset(max)
	for _, i := range arr {
		bs.Set(i)
	}

	id := 0
	for i := 0; i < max; i++ {
		if bs.IsSet(i) {
			arr[id] = i
			id++
		}
	}
}
