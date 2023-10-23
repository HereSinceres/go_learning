package loop_test

import (
	"fmt"
	"testing"
)

func TestWhilleLoop(t *testing.T) {
	n := 5
	for i := 0; i < n; i++ {
		fmt.Print(i)
	}
}
