package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringInit(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)
	t.Log(strings.Join(parts, "-"))

}

func TestStrinConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log(s)

	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
