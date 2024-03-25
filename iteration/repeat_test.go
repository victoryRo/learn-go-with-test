package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	exp := "aaaaa"
	actual := Repeat("a", 5)

	assert.Equal(t, exp, actual)
	if exp != actual {
		t.Errorf("expected %q but got %q", exp, actual)
	}
}

func TestCompare(t *testing.T) {
	data := map[int]struct {
		msg    string
		first  string
		second string
	}{
		0:  {msg: "return 0", first: "light", second: "light"},
		-1: {msg: "return -1", first: "light", second: "yellow"},
		1:  {msg: "return +1", first: "zero", second: "cero"},
	}

	for expected, param := range data {
		t.Run(param.msg, func(t *testing.T) {
			actual := Comparision(param.first, param.second)

			if expected != actual {
				t.Errorf("expected %d but got %d", expected, actual)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("c", 7)
	fmt.Println(res)

	//Output: ccccccc
}
