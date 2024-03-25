package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		exp := "Hello, Victor"
		actual := Hello("Victor", "")

		assertCorrectMessage(t, exp, actual)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		exp := "Hello, world"
		actual := Hello("", "")

		assertCorrectMessage(t, exp, actual)
	})
	t.Run("in Spanish", func(t *testing.T) {
		exp := "Hola, Maria"
		actual := Hello("Maria", "Spanish")

		assertCorrectMessage(t, exp, actual)
	})
	t.Run("in French", func(t *testing.T) {
		exp := "Bonjour, juana"
		actual := Hello("juana", "French")

		assertCorrectMessage(t, exp, actual)
	})
}

func assertCorrectMessage(t testing.TB, exp, actual string) {
	t.Helper()

	assert.Equal(t, exp, actual)
	if exp != actual {
		t.Errorf("expected %s but got %s", exp, actual)
	}
}
