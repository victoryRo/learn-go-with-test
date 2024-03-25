package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	exp := 4
	actual := Add(2, 2)

	assert.Equal(t, exp, actual)
	if exp != actual {
		t.Errorf("expected %d and got %d", exp, actual)
	}
}

func ExampleAdd() {
	sum := Add(4, 2)
	fmt.Println(sum)
	// Output: 6
}
