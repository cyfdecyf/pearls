package chap02

import (
	"github.com/cyfdecyf/goutil"
)

// RotateLeft rotates arr left by i position. Panics if i < 0.
func RotateLeft(arr goutil.SwapInterface, i int) {
	if i < 0 {
		panic("RotateLeft negative position")
	}
	n := arr.Len()
	goutil.Reverse(arr)
	goutil.ReverseRange(arr, 0, n-i)
	goutil.ReverseRange(arr, n-i, n)
}
