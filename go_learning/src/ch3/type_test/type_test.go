package typetest_test

import (
	"math"
	"testing"
)

// 不允许隐士类型转换

// 类型别名
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)

	var c MyInt

	c = MyInt(b)
	t.Log(a, b, c)
}

// 预定义值

func TestMath(t *testing.T) {
	t.Log(math.MaxInt)
}

// 指针类型
// 不支持指针运算

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 // 会报错，不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T, %T", a, aPtr)
}

// string 是值类型， 其默认的初始化值为空字符串 而不是nil
func TestString(t *testing.T) {

	var s string
	t.Log("*" + s + "*")

	t.Log(len(s))
	if s == "" {
		t.Log(len(s))

	}
}
