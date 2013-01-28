package chap02

import (
	"github.com/cyfdecyf/goutil/math"
)

// reverse reverse the arr in the range of [l,h)
func reverse(arr []int, l, h int) {
	n := h - l
	for i := 0; i < n/2; i++ {
		arr[l+i], arr[h-i-1] = arr[h-i-1], arr[l+i]
	}
}

// RotateLeft rotates arr left by i position. Panics if i < 0.
func RotateLeft(arr []int, i int) {
	if i < 0 {
		panic("RotateLeft: negative rotation distance")
	}
	n := len(arr)
	i = i % n
	reverse(arr, 0, n)
	reverse(arr, 0, n-i)
	reverse(arr, n-i, n)
}

// rotateLeftBlockSwap rotates item in the range of [l, h) by i position.
func rotateLeftBlockSwap(arr []int, l, h, i int) {
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
	ll := h - swapLen
	for j := 0; j < swapLen; j++ {
		arr[l+j], arr[ll+j] = arr[ll+j], arr[l+j]
	}
	if i == n-i {
		return
	}
	rotateLeftBlockSwap(arr, nl, nh, ni)
}

// RotateLeftBlockSwap rotates arr left by i position. It uses the recursive
// block swap algorithm. Panics if i > arr.Len() or i < 0.
func RotateLeftBlockSwap(arr []int, i int) {
	if i < 0 {
		panic("RotateLeftRecursive: negative rotation distance")
	}
	n := len(arr)
	i = i % n
	if i == 0 {
		return
	}
	rotateLeftBlockSwap(arr, 0, len(arr), i)
}

// RotateLeftJuggling rotates arr left by i position. It iteratively moves
// each item i position to the target place. Panics if i < 0.
func RotateLeftJuggling(arr []int, i int) {
	if i < 0 {
		panic("RotateLeftJuggling: negative rotation distance")
	}
	n := len(arr)
	i = i % n
	// Refer to Concrete Mathematics section 4.8
	// The following n numbers,
	// 0 mod n, i mod n, 2i mod n, ..., (n-1)i mod n
	// consist exact d=gcd(i, n) copies of the n/d numbers
	// 0, d, 2d, ..., m-d
	// in some order.
	// RotateLeft only needs the first copy and iterate d times plusing the
	// start index by one each time.
	// Items moved in each iteration are j, d+j, 2d+j, ..., m-d+j
	d := math.GCD(i, n)
	for j := 0; j < d; j++ {
		jval := arr[j]
		for p := j; ; {
			next := (p + i) % n
			if next == j {
				arr[p] = jval
				break
			}
			arr[p] = arr[next]
			p = next
		}
	}
}
