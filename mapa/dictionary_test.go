package mapa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		exp := "this is just a test"
		actual, _ := dictionary.Search("test")

		assertString(t, exp, actual)
	})
	t.Run("unkown word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")

		assertError(t, err, ErrNotFound)
	})
}

func ExampleDictionary_Search() {
	dictionary := Dictionary{"hi": "how are you"}
	actual, _ := dictionary.Search("hi")
	fmt.Println(actual)
	// Output: how are you
}

func TestAdd(t *testing.T) {
	key, value := "hi", "how are you"

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new text ok")

		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func ExampleDictionary_Add() {
	dictionary := Dictionary{}
	err := dictionary.Add("hi", "how are you")
	fmt.Println(err)
	// Output: <nil>
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key", func(t *testing.T) {
		key, value := "sunday", "first day, day of the sun"
		newvalue := "bright day"
		dictionary := Dictionary{key: value}

		err := dictionary.Update(key, newvalue)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newvalue)
	})
	t.Run("update with new key", func(t *testing.T) {
		key, value := "friday", "day of venus love"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "deleted"
	dictionary := Dictionary{key: "I will be removed"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", key)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	actual, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, value, actual)
}

func assertError(t testing.TB, actual, err error) {
	t.Helper()

	if actual != err {
		t.Errorf("actual error %q expected error %q", actual, err)
	}
}

func assertString(t testing.TB, exp, actual string) {
	t.Helper()

	assert.Equal(t, exp, actual)
	if exp != actual {
		t.Errorf("expected %q and got %q", exp, actual)
	}
}
