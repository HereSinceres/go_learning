package operatortest

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 23}
	b := [...]int{1, 23}
	// c := [...]int{1, 23, 2}

	t.Log(a == b)
	// 长度不同 不可以比较
	// t.Log(a == c)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
