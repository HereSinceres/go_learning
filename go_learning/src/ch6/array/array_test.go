package array_test

import "testing"

func TestXxx(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2}
	t.Log(arr, arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{12, 3, 4}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

// 数组截取

func TestArraySection(t *testing.T) {

	arr3 := [...]int{12, 3, 4, 5}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)
}

// 切片
