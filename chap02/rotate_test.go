package chap02

import (
	"github.com/cyfdecyf/goutil"
	"reflect"
	"testing"
)

func testRotateLeft(rotateFunc func(goutil.SwapInterface, int), msg string, t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	type arrRot struct {
		arr []int
		rot int
	}
	testData := []arrRot{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 0},
		{[]int{2, 3, 4, 5, 6, 7, 8, 1}, 7},
		{[]int{8, 1, 2, 3, 4, 5, 6, 7}, 1},
		{[]int{6, 7, 8, 1, 2, 3, 4, 5}, 3},
	}

	for _, td := range testData {
		rotateFunc(goutil.IntSwapSlice(td.arr), td.rot)
		if !reflect.DeepEqual(arr, td.arr) {
			t.Errorf("%s %d error, after rotation: %v", msg, td.rot, td.arr)
		}
	}
}

func TestRotateLeft(t *testing.T) {
	testRotateLeft(RotateLeft, "RotateLeft", t)
}

func TestRotateLeftRecursive(t *testing.T) {
	testRotateLeft(RotateLeftBlockSwap, "RotateLeftRecursive", t)
}
