package smi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, exp float64, shape Shape) {
		t.Helper()
		actual := shape.Area()

		assert.Equal(t, exp, actual)
		if exp != actual {
			t.Errorf("expected Area %.g and got %.g", exp, actual)
		}
	}

	data := []struct {
		msg   string
		shape Shape
		exp   float64
	}{
		{msg: "rectangle", shape: Rectangle{Width: 12.00, Heigth: 6.00}, exp: 72.0},
		{msg: "circle", shape: Circle{10}, exp: 314.1592653589793},
		{msg: "triangle", shape: Triangle{12, 6}, exp: 36},
	}

	for _, tt := range data {
		t.Run(tt.msg, func(t *testing.T) {
			checkArea(t, tt.exp, tt.shape)
		})
	}
}
