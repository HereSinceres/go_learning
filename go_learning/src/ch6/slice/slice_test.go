package slice

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)

	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0])
	s2 = append(s2, 2)
	t.Log(s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {

	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	t.Log(year, len(year), cap(year))
	Q2 := year[3:6]

	t.Log(Q2, len(Q2), cap(Q2))
}

// 数组固定大小 Slice 不固定

// 数组可以比较, Slice 不可以比较

func TestSliceComparing(t *testing.T) {

	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// if a == b {
	// 	t.Log("equal")
	// }
	t.Log(a, len(a), cap(b))

}
