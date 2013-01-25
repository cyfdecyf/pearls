package chap02

import (
	"github.com/cyfdecyf/goutil"
	"reflect"
	"testing"
)

func testRotateLeft(rotateFunc func(goutil.SwapInterface, int), msg string, t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rot7 := []int{2, 3, 4, 5, 6, 7, 8, 1}
	rot0 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rot1 := []int{8, 1, 2, 3, 4, 5, 6, 7}

	rotateFunc(goutil.IntSwapSlice(rot7), 7)
	if !reflect.DeepEqual(arr, rot7) {
		t.Errorf("%s 7 error, after rotation:", msg, rot7)
	}
	rotateFunc(goutil.IntSwapSlice(rot0), 0)
	if !reflect.DeepEqual(arr, rot0) {
		t.Errorf("%s 0 error, after rotation:", msg, rot0)
	}
	rotateFunc(goutil.IntSwapSlice(rot1), 1)
	if !reflect.DeepEqual(arr, rot1) {
		t.Errorf("%s 1 error, after rotation:", msg, rot1)
	}
}

func TestRotateLeft(t *testing.T) {
	testRotateLeft(RotateLeft, "RotateLeft", t)
}
