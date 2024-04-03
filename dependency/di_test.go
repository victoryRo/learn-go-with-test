package dependency

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Victoria")

	actual := buffer.String()
	exp := "Hello, Victoria"

	if exp != actual {
		t.Errorf("expected %s but actual is %s", exp, actual)
	}
}
