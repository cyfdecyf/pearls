package chap02

import (
	"fmt"
	"github.com/cyfdecyf/goutil"
)

var _ = fmt.Printf

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

// rotateLeftBlockSwap rotates item in the range of [l, h) by i position.
func rotateLeftBlockSwap(arr goutil.SwapInterface, l, h, i int) {
	n := h - l
	var nh, nl, ni, swapLen int
	if i <= n-i {
		// len(a) = len(b_r) = i
		// a b_l b_r -> b_r b_l a -> b_l b_r a
		swapLen = i
		nl = l
		nh = h - swapLen
		ni = swapLen
	} else {
		// len(b) = len(a_l) = n-i
		// a_l a_r b -> b a_r a_l -> b a_l a_r
		swapLen = n - i
		nl = l + swapLen
		nh = h
		ni = n - 2*swapLen
	}
	for j := 0; j < swapLen; j++ {
		arr.Swap(l+j, h-swapLen+j)
	}
	if i == n-i {
		return
	}
	rotateLeftBlockSwap(arr, nl, nh, ni)
}

// RotateLeftBlockSwap rotates arr left by i position. It uses the recursive
// block swap algorithm. Panics if i > arr.Len() or i < 0.
func RotateLeftBlockSwap(arr goutil.SwapInterface, i int) {
	if i < 0 || i > arr.Len() {
		panic("RotateLeftRecursive: rotation position invalid")
	}
	if i == 0 || i == arr.Len() {
		return
	}
	rotateLeftBlockSwap(arr, 0, arr.Len(), i)
}
