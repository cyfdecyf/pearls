package chap02

import (
	"reflect"
	"testing"
)

func testRotateLeft(rotateFunc func([]int, int), msg string, t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	type arrRot struct {
		arr []int
		rot int
	}
	testData := []arrRot{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 0},
		{[]int{8, 1, 2, 3, 4, 5, 6, 7}, 1},
		{[]int{7, 8, 1, 2, 3, 4, 5, 6}, 2},
		{[]int{6, 7, 8, 1, 2, 3, 4, 5}, 3},
		{[]int{5, 6, 7, 8, 1, 2, 3, 4}, 4},
		{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 5},
		{[]int{3, 4, 5, 6, 7, 8, 1, 2}, 6},
		{[]int{2, 3, 4, 5, 6, 7, 8, 1}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{[]int{6, 7, 8, 1, 2, 3, 4, 5}, 11},
	}

	for _, td := range testData {
		rotateFunc(td.arr, td.rot)
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

func TestRotateLeftJuggling(t *testing.T) {
	testRotateLeft(RotateLeftJuggling, "RotateLeftJuggling", t)
}

var benchArr = make([]int, 64*128) // 64 is common cache line size
var benchDis = 64

func BenchmarkRotateLeft_________(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateLeft(benchArr, benchDis)
	}
}

func BenchmarkRotateLeftBlockSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateLeftBlockSwap(benchArr, benchDis)
	}
}

func BenchmarkRotateLeftJuggling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateLeftJuggling(benchArr, benchDis)
	}
}
