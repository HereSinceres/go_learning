package consstant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Open = 1 << iota
	Close
	Pending
)

func TestDay(t *testing.T) {
	const a = iota
	t.Log(a)
	t.Log(Monday, Tuesday)
	t.Log(Open, Close, Pending)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestTry(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
