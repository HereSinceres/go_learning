package map1

import "testing"

func TestInitMap(t *testing.T) {

	m1 := map[int]int{1: 1}

	t.Log(m1[1])
	t.Logf("%d", len(m1))

	m2 := map[int]int{}
	m2[3] = 3

	t.Log(m2[1])
	t.Logf("%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("%d", len(m3))
}

func TestAccessNotExistingKy(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])

	if v, ok := m1[3]; ok {
		t.Log(v)
	} else {
		t.Log("not existing")

	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 3}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
