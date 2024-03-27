package array

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		values := []int{3, 4, 7, 8, 1, 2}

		exp := 25
		actual := Sum(values)

		assert.Equal(t, exp, actual)
		if exp != actual {
			t.Fatalf("expected %d But got %d GIVEN %v", exp, actual, values)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t testing.TB, exp, actual []int) {
		t.Helper()

		assert.Equal(t, exp, actual)
		if !reflect.DeepEqual(exp, actual) {
			t.Errorf("expected %v and got %v", exp, actual)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		exp := []int{3, 1}
		actual := SumAllTails([]int{4, 3}, []int{3, 1})

		checkSum(t, exp, actual)
	})
	t.Run("safely sum of empty slice", func(t *testing.T) {
		exp := []int{0, 7}
		actual := SumAllTails([]int{}, []int{3, 4, 2, 1})

		checkSum(t, exp, actual)
	})
}
